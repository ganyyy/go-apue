// +build ignore

package main

import (
	"log"
	"os"
	"syscall"
)

const (
	SEEK_SET = 0 // 绝对偏移量, 等同于从文件开头开始
	SEEK_CUR = 1 // 在当前位置上进行相对偏移
	SEEK_END = 2 // 从文件末尾进行绝对偏移
)

func main() {
	var err error
	_, err = syscall.Seek(int(os.Stdin.Fd()), 0, SEEK_CUR)
	if err != nil {
		log.Printf("cannot seek:%v", err)
	} else {
		log.Printf("seek success")
	}
}
