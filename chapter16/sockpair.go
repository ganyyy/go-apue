// +build ignore

package main

import (
	"go-apue/helper"
	"log"
	"sync"
	"syscall"
)

func main() {

	var fd, err = syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	helper.PanicIfError("socketpair", err)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		var buf [128]byte
		var n, _ = syscall.Read(fd[0], buf[:])
		log.Println(string(buf[:n]))
	}()

	syscall.Write(fd[1], []byte("hello world 12344"))

	waitGroup.Wait()
}
