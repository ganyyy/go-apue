// +build ignore

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

func pr_ids(name string) {

	var tty int
	syscall.RawSyscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGPGRP), uintptr(unsafe.Pointer(&tty)))

	log.Printf("%s, pid:%v, ppid:%v, pgrp:%v, tpgrp:%v", name, syscall.Getpid(), syscall.Getppid(), syscall.Getpgrp(), tty)
}

func main() {
	var pid uintptr
	var err error
	var errno syscall.Errno
	var noError = syscall.Errno(0)

	pr_ids("parent")

	if pid, _, errno = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0); errno != noError {
		panic(errno)
	} else if pid > 0 {
		time.Sleep(time.Second * 5)
	} else {
		pr_ids("child")
		var sigChan = make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGHUP)
		go func() {
			var sig = <-sigChan
			log.Printf(sig.String())
		}()

		err = syscall.Kill(syscall.Getpid(), syscall.SIGTSTP) // 暂停等待唤醒, 父进程结束后会触发 SIGCONT, 但是一直不成功, why?
		if err != nil {
			log.Printf(err.Error())
			os.Exit(-1)
		}

		pr_ids("child")

		if n, e := syscall.Read(syscall.Stdin, []byte{0}); n != 1 || e != nil {
			log.Printf("read error %v on controlling TTY", e)
		}
	}

	os.Exit(0)
}
