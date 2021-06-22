// +build ignore

package main

import (
	"go-apue/helper"
	"syscall"
	"time"
)

func main() {
	var pid uintptr
	var err error

	pid, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)

	if err != syscall.Errno(0) {
		helper.PanicIfError("fork1", err)
	}

	if pid != 0 {
		time.Sleep(time.Second * 2)
		syscall.Exit(2)
	}

	pid, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != syscall.Errno(0) {
		helper.PanicIfError("fork2", err)
	}
	if pid != 0 {
		err = syscall.Exec("/bin/dd", []string{"dd", "if=/etc/passwd", "of=/dev/null"}, []string{})
		helper.PanicIfError("dd", err)
		syscall.Exit(7)
	}

	pid, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != syscall.Errno(0) {
		helper.PanicIfError("fork3", err)
	}
	if pid != 0 {
		time.Sleep(time.Second * 8)
		syscall.Exit(0)
	}

	time.Sleep(time.Second * 6)
	syscall.Kill(syscall.Getpid(), syscall.SIGKILL)
	syscall.Exit(6)

}
