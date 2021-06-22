// +build ignore

package main

import (
	"fmt"
	"go-apue/helper"
	"syscall"
	"time"
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
		pid, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

		if err != syscall.Errno(0) {
			helper.PanicIfError("fork1", err)
		}
		if pid != 0 {
			// 子进程退出
			syscall.Exit(0)
		}
		// 子进程的子进程先睡一会
		time.Sleep(time.Second * 10)
		// 父进程(第一个子进程)退出后, 子进程(子进程的子进程)由`init`进程接管
		fmt.Printf("child child, parent pid = %v\n", syscall.Getppid())
		syscall.Exit(0)
	}

	var status syscall.WaitStatus
	var rusage syscall.Rusage
	var wpid int
	var err2 error
	// 等待子进程
	wpid, err2 = syscall.Wait4(int(pid), &status, syscall.WCONTINUED|syscall.WUNTRACED, &rusage)

	if wpid != int(pid) || err2 != nil {
		helper.PanicIfError("wait4", fmt.Errorf("pid:%v, wpid:%v, error:%v", pid, wpid, err2))
	}
	fmt.Printf("%+v, %+v", status, rusage)
	syscall.Exit(0)
}
