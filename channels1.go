package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func(){
		for i := 0; i<10; i++ {
			c <- i //blocking nature of unbuffered channels
			// this code stops until something take the value from channel
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}