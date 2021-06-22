// +build ignore

package main

import (
	"syscall"
)

func main() {
	var pid uintptr

	pid, _, _ = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)

	if pid == 0 {
		syscall.Exec("/home/gan/code/go/go-apue/chapter8/echoall", []string{"echoall", "argv1", "argv2"}, []string{"123", "456"})
	}

	syscall.Wait4(int(pid), nil, syscall.WCONTINUED|syscall.WUNTRACED, nil)

	pid, _, _ = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)

	if pid == 0 {
		syscall.Exec("echoall", []string{"echoall", "argv1"}, []string{"/home/gan/code/go/go-apue/chapter8/"})
	}

	syscall.Wait4(int(pid), nil, syscall.WCONTINUED|syscall.WUNTRACED, nil)
}
