// +build ignore

package main

import (
	"log"
	"syscall"
)

func main() {
	var stat syscall.Stat_t

	if err := syscall.Stat("foo", &stat); err != nil {
		log.Printf("stat foo error:%v", err)
	}

	if err := syscall.Chmod("foo", (stat.Mode|syscall.S_IXGRP)|syscall.S_ISGID|syscall.S_ISUID); err != nil {
		log.Printf("Chmod foo error:%v", err)
	}

	if err := syscall.Chmod("bar", syscall.S_IRUSR|syscall.S_IWUSR|syscall.S_IRGRP|syscall.S_IROTH); err != nil {
		log.Printf("Chmod bar error:%v", err)
	}
}
