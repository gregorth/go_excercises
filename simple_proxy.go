package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// header implements flag.Value interface.
type header struct {
	http.Header
}

func (h *header) String() string {
	return fmt.Sprintf("%v", h.Header)
}

func (h *header) Set(value string) error {
	if h.Header == nil {
		h.Header = make(http.Header)
	}
	kv := strings.SplitN(value, ":", 2)
	var k, v string
	k = strings.TrimSpace(kv[0])
	if len(kv) == 2 {
		v = strings.TrimSpace(kv[1])
	}
	h.Add(k, v)
	return nil
}

var extraHeader header

var server = &ProxyServer{
	Logger: log.New(os.Stderr, "", log.LstdFlags),
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: %s <PROXY_TO>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.StringVar(&server.Addr, "addr", ":8080", "bind address")
	flag.Var(&extraHeader, "header", `custom headers "Name:Value" (multiple)`)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	proxyTo, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	server.ProxyTo = proxyTo
	server.ExtraHeader = extraHeader.Header

	errCh := make(chan error)
	go func() {
		if err := server.Serve(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()
	log.Printf("Starting proxy server at %s\n", server.Addr)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	select {
	case err, ok := <-errCh:
		if ok {
			log.Println("Proxy server error:", err)
		}

	case sig := <-sigCh:
		log.Printf("Signal %s received\n", sig)
		if err := server.Shutdown(); err != nil {
			log.Println("Failed to shutdown proxy server:", err)
		}
		log.Println("Proxy server shutdown")
	}

}

type ProxyServer struct {
	Addr        string
	ProxyTo     *url.URL
	ExtraHeader http.Header
	Logger      *log.Logger
	MustClose   bool
	server      *http.Server
}

func (p *ProxyServer) Serve() error {
	proxy := httputil.NewSingleHostReverseProxy(p.ProxyTo)
	origDirector := proxy.Director
	proxy.Director = func(r *http.Request) {
		origURL := r.URL.String()
		origDirector(r)
		for name, values := range p.ExtraHeader {
			for _, value := range values {
				r.Header.Add(name, value)
			}
		}
		r.Host = r.URL.Host
		if p.Logger != nil {
			p.Logger.Printf("%s -> %s\n", origURL, r.URL)
		}
	}

	p.server = &http.Server{
		Addr:    p.Addr,
		Handler: proxy,
	}

	if err := p.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}

	return nil
}

func (p *ProxyServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.ShutdownWithContext(ctx)
}

func (p *ProxyServer) ShutdownWithContext(ctx context.Context) error {
	if err := p.server.Shutdown(ctx); err != nil {
		if p.MustClose {
			p.server.Close()
		}
		return err
	}

	return nil
}
