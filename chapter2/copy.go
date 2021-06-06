// +build ignore

package main

import (
	"fmt"
	"go-apue/helper"
	"os"
	"syscall"
)

//go:generate dd if=/dev/zero of=readme.data bs=100M count=2
//go:generate echo done

func main() {

	if len(os.Args) != 3 {
		helper.PanicIfError("args", fmt.Errorf("must input read/write file. cur args: %v", os.Args[1:]))
	}

	var pipe [2]int
	var err error
	err = syscall.Pipe(pipe[:])
	helper.PanicIfError("pipe", err)

	var writeFd, readFd int
	var writePath, readPath = os.Args[1], os.Args[2]

	writeFd, err = syscall.Open(writePath, syscall.O_WRONLY|syscall.O_CREAT, 0644)
	helper.PanicIfError("open write file", err)

	readFd, err = syscall.Open(readPath, syscall.O_RDONLY, 0)
	helper.PanicIfError("open read file", err)

	const (
		BufLen = 4096
	)
	const (
		SPLICE_F_MOVE = 1 << iota
		SPLICE_F_NONBLOCK
		SPLICE_F_MORE
		SPLICE_F_GIFT
	)
	var readCount, writeCount int64
	var fileInfo syscall.Stat_t
	err = syscall.Stat(readPath, &fileInfo)

	helper.PanicIfError("stat file", err)
	var n int64
	var total = fileInfo.Size
	for total > 0 {
		_, err = syscall.Splice(readFd, &readCount, pipe[1], nil, BufLen, SPLICE_F_MORE|SPLICE_F_MOVE)
		// if err != nil {
		// 	break
		// }
		n, err = syscall.Splice(pipe[0], nil, writeFd, &writeCount, BufLen, SPLICE_F_MORE|SPLICE_F_MOVE)

		total -= n
	}

}
