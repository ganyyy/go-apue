// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	var stat syscall.Stat_t
	var err error
	var fType string

	var fIsType = func(mode, typ uint32) bool {
		return mode&syscall.S_IFMT == typ
	}

	for _, arg := range os.Args[1:] {
		err = syscall.Lstat(arg, &stat)
		if err != nil {
			log.Printf("%v stat error:%v", arg, err)
			continue
		}

		switch {
		case fIsType(stat.Mode, syscall.S_IFREG):
			fType = "regular"
		case fIsType(stat.Mode, syscall.S_IFDIR):
			fType = "directory"
		case fIsType(stat.Mode, syscall.S_IFCHR):
			fType = "character special"
		case fIsType(stat.Mode, syscall.S_IFBLK):
			fType = "block special"
		case fIsType(stat.Mode, syscall.S_IFIFO):
			fType = "fifo"
		case fIsType(stat.Mode, syscall.S_IFLNK):
			fType = "symbolic link"
		case fIsType(stat.Mode, syscall.S_IFSOCK):
			fType = "socket"
		}

		fmt.Printf("%s type is %s\n", arg, fType)
	}
}
