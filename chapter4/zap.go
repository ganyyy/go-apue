// +build ignore

package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	var fd int
	var stat syscall.Stat_t
	var timespec [2]syscall.Timeval
	var err error
	for _, path := range os.Args[1:] {
		if err = syscall.Stat(path, &stat); err != nil {
			log.Printf("%v state error:%v", path, err)
			continue
		}
		if fd, err = syscall.Open(path, syscall.O_TRUNC|syscall.O_RDWR, 0); err != nil {
			log.Printf("%v open error:%v", path, err)
			continue
		}

		timespec[0].Sec = stat.Atim.Sec
		timespec[0].Usec = stat.Atim.Nsec / 1000
		timespec[1].Sec = stat.Mtim.Sec
		timespec[1].Usec = stat.Mtim.Nsec / 1000
		if err = syscall.Futimes(fd, timespec[:]); err != nil {
			log.Printf("%v Futimes error:%v", path, err)
			continue
		}
		log.Printf("%v a_t:%v, m_t:%v", path, timespec[0], timespec[1])
	}
}
