// +build ignore

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var sigChan = make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGCONT)

	go func() {
		for sig := range sigChan {
			switch sig {
			case syscall.SIGUSR1, syscall.SIGUSR2:
				log.Printf("receive:%v", sig)
			case syscall.SIGCONT:
				log.Printf("receive continue:%v", sig)
			default:
				log.Printf("other sig:%v", sig)
			}
		}
	}()

	// if !signal.Ignored(syscall.SIGINT) {
	// 	signal.Ignore(syscall.SIGINT)
	// }
	// if !signal.Ignored(syscall.SIGQUIT) {
	// 	signal.Ignore(syscall.SIGQUIT)
	// }

	select {}
}
