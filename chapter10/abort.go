// +build ignore

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var sigChan = make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGABRT)

	go func() {
		for sig := range sigChan {
			log.Println(sig)
		}
	}()

	time.Sleep(time.Second * 20)
}
