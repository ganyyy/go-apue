// +build ignore

package main

import (
	"fmt"
	"go-apue/helper"
	"syscall"
)

func main() {
	var pid uintptr
	var err syscall.Errno

	pid, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != syscall.Errno(0) {
		helper.PanicIfError("fork1", err)
	}

	if pid == 0 {

		// 子进程
		var s = "this is child\n"
		for i := 0; i < len(s); i++ {
			fmt.Print(string(s[i]))
		}
	} else {
		// 父进程

		var s = "this is parent\n"
		for i := 0; i < len(s); i++ {
			fmt.Print(string(s[i]))
		}
	}
}
