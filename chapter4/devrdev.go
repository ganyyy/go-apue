// +build ignore

package main

import (
	"go-apue/helper"
	"log"
	"os"
	"syscall"
)

func main() {
	var stat syscall.Stat_t
	var err error

	if err = syscall.Stat(os.Args[1], &stat); err != nil {
		helper.PanicIfError("stat", err)
	}

	log.Println(stat.Dev, stat.Rdev)
}
