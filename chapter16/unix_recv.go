// +build ignore

package main

import (
	"go-apue/helper"
	"log"
	"net"
	"os"
	"syscall"
)

func main() {
	syscall.Unlink(helper.PATH)
	var laddr *net.UnixAddr
	var err error
	laddr, err = net.ResolveUnixAddr("unix", helper.PATH)
	helper.PanicIfError("resolve local addr", err)

	var listener *net.UnixListener
	listener, err = net.ListenUnix("unix", laddr)
	helper.PanicIfError("listen error", err)

	var conn *net.UnixConn

	conn, err = listener.AcceptUnix()
	helper.PanicIfError("accept unix", err)
	var connFd int
	var rawConn syscall.RawConn
	rawConn, err = conn.SyscallConn()
	helper.PanicIfError("raw conn", err)
	rawConn.Control(func(fd uintptr) {
		connFd = int(fd)
	})

	var fd int
	var msg string
	fd, msg, err = helper.RecvFD(connFd)
	helper.PanicIfError("recvFD", err)
	log.Println("recv msg:%v", msg)
	var f *os.File
	f = os.NewFile(uintptr(fd), "")
	defer f.Close()
	var buf [1024]byte
	var n int
	n, err = f.Read(buf[:])
	helper.PanicIfError("read", err)
	log.Println(string(buf[:n]))

}
