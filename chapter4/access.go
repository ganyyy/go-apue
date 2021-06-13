// +build ignore

package main

import (
	"fmt"
	"go-apue/helper"
	"log"
	"os"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		helper.PanicIfError("check gs", fmt.Errorf("input path"))
	}

	var path = os.Args[1]

	var err error

	const (
		R_OK = 4
		W_OK = 2
		X_OK = 1
		F_OK = 0
	)

	if err = syscall.Access(path, R_OK); err != nil {
		log.Printf("access error:%v", err)
	} else {
		log.Printf("access ok")
	}

	if _, err = syscall.Open(path, syscall.O_RDONLY, 0); err != nil {
		log.Printf("open error:%v", err)
	} else {
		log.Printf("open ok")
	}

}
