// +build ignore

package main

import (
	"log"
	"syscall"
)

const (
	SEEK_SET = 0 // 绝对偏移量, 等同于从文件开头开始
	SEEK_CUR = 1 // 在当前位置上进行相对偏移
	SEEK_END = 2 // 从文件末尾进行绝对偏移
)

//go:generate od -c file.hole

func main() {
	var (
		buf1 = []byte("123456")
		buf2 = []byte("abcdef")
	)

	var fd int
	var err error

	if fd, err = syscall.Creat("file.hole", 0644); err != nil {
		log.Panicf("creat error:%v", err)
	}

	if _, err = syscall.Write(fd, buf1); err != nil {
		log.Panicf("write1 error:%v", err)
	}

	if _, err = syscall.Seek(fd, 1000, SEEK_SET); err != nil {
		log.Panicf("seek error:%v", err)
	}

	if _, err = syscall.Write(fd, buf2); err != nil {
		log.Printf("write2 error:%v", err)
	}

}
