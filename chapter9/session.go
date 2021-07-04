// +build ignore

package main

import (
	"syscall"
)

func main() {

	var pid uintptr
	var errNo syscall.Errno
	var err error

	if pid, _, errNo = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0); errNo != syscall.Errno(0) {
		panic(errNo)
	}

	if pid == 0 {
		// 子进程
		var npid int
		npid, err = syscall.Setsid()

		syscall.ForkExec()

	} else {
		// 父进程
	}
}
