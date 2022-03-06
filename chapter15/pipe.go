// +build ignore

package main

import (
	"log"
	"syscall"
)

func main() {
	var n int
	var fd [2]int
	var pid uintptr
	var err error

	if err = syscall.Pipe(fd[:]); err != nil {
		log.Panicf("pipe error:%v", err)
	}

	pid, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != syscall.Errno(0) {
		log.Panicf("fork error:%v", err)
	}

	if pid > 0 {
		// 父进程
		syscall.Close(fd[0]) // 关闭读端
		syscall.Write(fd[1], []byte("hello world"))
	} else {
		// 子进程
		syscall.Close(fd[1]) // 关闭写端
		var buf [128]byte
		n, _ = syscall.Read(fd[0], buf[:])
		syscall.Write(syscall.Stdout, buf[:n])
	}
	syscall.Exit(0)
}
