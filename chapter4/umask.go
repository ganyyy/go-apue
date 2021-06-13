// +build ignore

package main

import "syscall"

func main() {
	const (
		RWRWRW = syscall.S_IRUSR | syscall.S_IWUSR | syscall.S_IRGRP | syscall.S_IWGRP | syscall.S_IROTH | syscall.S_IWOTH
	)

	syscall.Umask(0)
	syscall.Creat("foo", RWRWRW)

	syscall.Umask(syscall.S_IRGRP | syscall.S_IWGRP | syscall.S_IROTH | syscall.S_IWOTH)
	syscall.Creat("bar", RWRWRW)

	syscall.Umask(0)
	syscall.Creat("test", syscall.S_IRUSR|syscall.S_IXUSR|syscall.S_IRGRP|syscall.S_IXOTH|syscall.S_IROTH|syscall.S_ISGID|syscall.S_ISUID|syscall.S_ISVTX)
}
