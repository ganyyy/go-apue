// +build ignore

package main

import (
	"go-apue/helper"
	"log"
	"net"
	"sync"
	"syscall"
	"time"
)

func main() {
	const NAME = "/tmp/test_unix_sock"
	// 防止已存在
	syscall.Unlink(NAME)
	var addr net.UnixAddr
	addr.Name = NAME
	var listener, err = net.ListenUnix("unix", &addr)
	defer syscall.Unlink(NAME)
	helper.PanicIfError("listen unix", err)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		var conn, err = listener.AcceptUnix()
		helper.PanicIfError("accept  unix", err)
		conn.Write([]byte("hello world"))
		wg.Done()
	}()

	var conn *net.UnixConn
	conn, err = net.DialUnix("unix", nil, &addr)
	helper.PanicIfError("dial unix", err)

	var buf [128]byte
	var n int
	n, err = conn.Read(buf[:])
	helper.PanicIfError("read", err)
	log.Println(string(buf[:n]))

	time.Sleep(time.Second * 10)
	wg.Wait()
}
