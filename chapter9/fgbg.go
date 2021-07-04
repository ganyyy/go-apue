// +build ignore

package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	var sigChan = make(chan os.Signal, 5)

	signal.Notify(sigChan)

	go func() {
		for {
			select {
			case sig := <-sigChan:
				log.Print(sig)
			}
		}
	}()

	time.Sleep(time.Second * 20)
}
