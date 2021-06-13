// +build ignore

package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println(os.Getwd())
	os.Chdir("/tmp")
	fmt.Println(os.Getwd())

	fmt.Println(syscall.PathMax,
		syscall.NAME_MAX)
}
