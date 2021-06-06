// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	var pathFd int
	var err error
	// 打开一个目录文件描述符, 该描述符必须是只读形式打开(目录不可写)
	pathFd, err = syscall.Open("/home/gan/code/go/go-apue/chapter2", syscall.O_RDONLY|syscall.O_DIRECTORY, 0)

	if err != nil {
		log.Printf("open error:%v", err)
		os.Exit(1)
	}

	var fileFd int
	// 基于指定的目录fd打开
	fileFd, err = syscall.Openat(pathFd, "Makefile", syscall.O_RDONLY, 0)
	if err != nil {
		log.Printf("open file: %v", err)
		os.Exit(1)
	}

	var buf [1024]byte
	var n int
	n, err = syscall.Read(fileFd, buf[:])
	if err != nil {
		log.Printf("Read file: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(buf[:n]))

	var fd2 int
	// 在执行路径下打开
	const (
		AT_FDCWD = -100
	)
	fd2, err = syscall.Openat(AT_FDCWD, "test.txt", syscall.O_RDWR|syscall.O_CREAT, 0664)
	if err != nil {
		log.Printf("open write file: %v", err)
		os.Exit(1)
	}

	n, err = syscall.Write(fd2, buf[:n])
	if err != nil {
		log.Printf("write file: %v", err)
		os.Exit(1)
	}
}
