// +build ignore

package main

import (
	"log"
	"os"
	"syscall"
)

const (
	DEF_PAGER = "/bin/more"
)

func main() {
	var n, n2 int
	var pipe [2]int
	var pid uintptr
	var buf [1024]byte
	var fd int
	var err error

	if len(os.Args) != 2 {
		log.Panicf("please input src file")
	}

	fd, err = syscall.Open(os.Args[len(os.Args)-1], syscall.O_RDONLY, 0)
	if err != nil {
		log.Panicf("cannot read %s, %v", os.Args[len(os.Args)-1], err)
	}

	err = syscall.Pipe(pipe[:])
	if err != nil {
		log.Panicf("pipe error:%v", err)
	}

	pid, _, err = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != syscall.Errno(0) {
		log.Panicf("fork error:%v", err)
	}

	if pid > 0 {
		// 读取文件并写入到 管道的写端
		for n, err = syscall.Read(fd, buf[:]); n != 0 && err == nil; n, err = syscall.Read(fd, buf[:]) {
			if n2, err = syscall.Write(pipe[1], buf[:n]); n2 != n {
				log.Panicf("write to pipe error")
			}
		}

		syscall.Close(pipe[1])

		syscall.Wait4(int(pid), nil, 0, nil)
		os.Exit(0)
	} else {
		// 关闭子进程中的写端
		syscall.Close(pipe[1])
		if pipe[0] != syscall.Stdin {
			// 将 管道的读端复制到 标准输入. 作为接下来的 PAGER 程序的输入
			err = syscall.Dup2(pipe[0], syscall.Stdin)
			if err != nil {
				log.Panicf("dup2 error:%v", err)
			}
		}

		var pager = os.Getenv("PAGER")
		// if len(pager) == 0 {
		pager = DEF_PAGER
		// }

		err = syscall.Exec(pager, nil, nil)
		if err != nil {
			log.Panicf("exec error:%v", err)
		}
	}

}
