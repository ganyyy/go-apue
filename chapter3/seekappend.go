// +build ignore

package main

import (
	"go-apue/helper"
	"log"
	"syscall"
)

const (
	SEEK_SET = 0 // 绝对偏移量, 等同于从文件开头开始
	SEEK_CUR = 1 // 在当前位置上进行相对偏移
	SEEK_END = 2 // 从文件末尾进行绝对偏移
)

//go:generate echo "12345" > ./a.txt

func main() {
	var fd int
	var err error
	var off int64
	var n int

	fd, err = syscall.Open("a.txt", syscall.O_RDWR|syscall.O_APPEND, 0644)
	helper.PanicIfError("open file", err)

	off, err = syscall.Seek(fd, 0, SEEK_SET)
	helper.PanicIfError("seek", err)
	_ = off
	var buf [10]byte
	n, err = syscall.Read(fd, buf[:])
	helper.PanicIfError("read", err)
	log.Printf("read: %v", buf[:n])

	n, err = syscall.Write(fd, buf[:n])
	helper.PanicIfError("write", err)
}
