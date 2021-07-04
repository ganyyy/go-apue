// +build ignore

package main

import (
	"log"
	"syscall"
	"time"
)

func main() {
	// var sigChan = make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGALRM)
	// go func() {

	// 	for sig := range sigChan {
	// 		log.Printf("receive sig:%v", sig)
	// 	}
	// }()

	var tick = time.NewTicker(time.Second * 2)

	for {
		select {
		case <-tick.C:
			var _, _, err = syscall.Syscall(syscall.SYS_ALARM, 1, 0, 0)
			log.Printf("alarm error:%v", err)
		}

	}

}
