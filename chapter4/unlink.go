// +build ignore

package main

import (
	"go-apue/helper"
	"log"
	"syscall"
	"time"
)

func main() {
	var fd, err = syscall.Open("RUNNING", syscall.O_CREAT|syscall.O_APPEND, 0644)
	helper.PanicIfError("create tmp file", err)
	time.Sleep(time.Second * 5)
	err = syscall.Unlink("RUNNING")
	helper.PanicIfError("unlink file", err)

	_, err = syscall.Write(fd, []byte("RUNNING"))
	if err != nil {
		log.Printf("write log:%v", err)
	}
	time.Sleep(time.Second * 15)
}
