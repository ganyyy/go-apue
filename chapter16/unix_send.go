// +build ignore

package main

import (
	"go-apue/helper"
	"net"
	"os"
	"syscall"
)

func main() {
	var raddr *net.UnixAddr
	var err error
	var conn *net.UnixConn
	raddr, err = net.ResolveUnixAddr("unix", helper.PATH)
	helper.PanicIfError("resolve remote addr", err)
	conn, err = net.DialUnix("unix", nil, raddr)
	helper.PanicIfError("dial unix", err)

	var file *os.File
	file, err = os.Open("./unix.go")
	helper.PanicIfError("open file", err)
	defer file.Close()
	var sysConn syscall.RawConn
	sysConn, err = conn.SyscallConn()
	helper.PanicIfError("SyscallConn", err)

	// 获取原始的FD套接字
	var connFd int
	sysConn.Control(func(fd uintptr) {
		connFd = int(fd)
	})

	err = helper.SendFD(connFd, int(file.Fd()), "hello world")
	helper.PanicIfError("send fd", err)

}
