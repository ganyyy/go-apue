// +build ignore

package main

import (
	"log"
	"syscall"
	"unsafe"
)

func main() {
	var attr syscall.Termios
	r1, r2, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TCGETS, uintptr(unsafe.Pointer(&attr)))
	log.Printf("r1:%v, r2:%v, error:%v, val:%+v", r1, r2, err, attr)
}
