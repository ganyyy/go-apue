// +build ignore

package main

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

var global = 200

func main() {
	var val = 100

	syscall.Write(syscall.Stdout, []byte("123456\n"))

	// Golang的标准库输出貌似是无缓冲的?
	fmt.Printf("before fork")
	var pid, _, err = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if err != syscall.Errno(0) {
		log.Printf("error:%v", err)
	}
	if pid != 0 {
		time.Sleep(time.Second)
	} else {
		global++
		val++
	}

	fmt.Printf("pid:%v, glob:%v, var:%v\n", pid, global, val)
}
