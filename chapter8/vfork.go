// +build ignore

package main

import (
	"fmt"
	"log"
	"syscall"
)

var global = 200

func main() {
	var val = 100

	syscall.Write(syscall.Stdout, []byte("123456\n"))

	// Golang的标准库输出貌似是无缓冲的?
	fmt.Printf("before fork\n")
	var pid, _, err = syscall.RawSyscall(syscall.SYS_VFORK, 0, 0, 0)
	if err != syscall.Errno(0) {
		log.Printf("error:%v", err)
	}
	if pid == 0 {
		global++
		val++
		fmt.Printf("child pid:%v, glob:%v, var:%v\n", pid, global, val)
		// vfork必须要通过`exec/exit`结束
		syscall.Exit(1)
	}

	fmt.Printf("parent pid:%v, glob:%v, var:%v\n", pid, global, val)
	syscall.Exit(0)
}
