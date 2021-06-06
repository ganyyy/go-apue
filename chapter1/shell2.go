// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

func main() {
	var s string
	var err error

	go func() {
		var sigChan = make(chan os.Signal, 1)
		signal.Notify(sigChan)
		for {
			var c = <-sigChan
			switch c {
			case os.Interrupt:
				log.Println("receive EINTR")
				os.Exit(-1)
			default:
				log.Printf("receive %v", c)
			}
		}
	}()

	for _, err = fmt.Scanf("%s", &s); err == nil; _, err = fmt.Scanf("%s", &s) {
		var cmd = exec.Command(s)
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			log.Printf("exec %v error %v", s, err)
		}
		_ = cmd.Wait()
	}

	log.Println("shell exit")
}
