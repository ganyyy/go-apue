// +build ignore

package main

import (
	"go-apue/helper"
	"syscall"
)

func main() {
	helper.PanicIfError("acct", syscall.Acct("./test1"))
}
