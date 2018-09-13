package main

import (
	"net/http"
	"fmt"
	"time"
	"net"
	"log"
	"os"
)

type tcpKeepAliveListener struct {
	*net.TCPListener
}
func myHandler(w http.ResponseWriter, r *http.Request){
	message := "Hello world"
	w.Write([]byte(message))
	fmt.Println("request served: ")
}

//var s = &http.Server{
//	Addr:           ":8888",
//	Handler:        nil,
//	ReadTimeout:    10 * time.Second,
//	WriteTimeout:   10 * time.Second,
//	MaxHeaderBytes: 1 << 20,
//}
type GServer struct {
	Addr	string
	Handler http.Handler
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int

}
func (srv *GServer) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve.tcpKeepAliveListener{ln.(*net.TCPListener)})
}

func (srv *GServer) Serve(l net.Listener) error {

}

var http_server = &GServer{}

func main() {
	//write simple server that with each request print 'foo' into stdout



	http_server.ListenAndServe()
	//err := http.ListenAndServe(":8888", nil)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//}
}
