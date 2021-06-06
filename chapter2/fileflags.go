// +build ignore

package main

import (
	"fmt"
	"go-apue/helper"
	"log"
	"os"
	"strconv"
	"syscall"
)

//go:generate make fileflags
//go:generate ./fileflags 0 < /dev/tty
//go:generate ./fileflags 1 > /tmp/tmp
//go:generate ./fileflags 2 2>>/tmp/tmp
//go:generate ./fileflags 5 5<>/tmp/tmp

func main() {
	var val uintptr
	if len(os.Args) != 2 {
		helper.PanicIfError("args error", fmt.Errorf("must need a file path"))
	}
	var err error
	var fd int
	fd, err = strconv.Atoi(os.Args[1])
	helper.PanicIfError("atoi", err)

	val, _, err = syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd), syscall.F_GETFL, 0)
	if err != syscall.Errno(0) {
		helper.PanicIfError("fcntl.F_GETFL", err)
	}

	switch val & syscall.O_ACCMODE {
	case syscall.O_RDONLY:
		log.Println("read only")
	case syscall.O_WRONLY:
		log.Println("write only")
	case syscall.O_RDWR:
		log.Println("rdwr only")
	default:
		log.Println("unknown access mode")
	}

	if val&syscall.O_APPEND != 0 {
		log.Println("O_APPEND")
	}
	if val&syscall.O_NONBLOCK != 0 {
		log.Println("O_NONBLOCK")
	}
	if val&syscall.O_SYNC != 0 {
		log.Println("O_SYNC")
	}
	if val&syscall.O_FSYNC != 0 {
		log.Println("O_FSYNC")
	}
}
