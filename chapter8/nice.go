// +build ignore

package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

func main() {

	var ts, et syscall.Tms
	var stt, ett uintptr
	var err error
	stt, err = syscall.Times(&ts)
	if err != nil {
		log.Printf("times error:%v", err)
	}

	var prio, _ = syscall.Getpriority(syscall.PRIO_PROCESS, 0)
	fmt.Printf("process prio:%v\n", prio)

	syscall.Setpriority(syscall.PRIO_PROCESS, 0, 50)
	prio, _ = syscall.Getpriority(syscall.PRIO_PROCESS, 0)
	fmt.Printf("process prio:%v\n", prio)

	time.Sleep(time.Second * 2)

	ett, err = syscall.Times(&et)
	if err != nil {
		log.Printf("times error:%v", err)
	}

	fmt.Printf("tick:%v, start ts:%#v, end ts:%#v", ett-stt, ts, et)

}
