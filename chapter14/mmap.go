// +build ignore

package main

import (
	"log"
	"syscall"
)

func main() {
	var fd int
	var err error

	fd, err = syscall.Open("./file.txt", syscall.O_RDONLY, 0)
	if err != nil {
		log.Fatalln("open", err)
	}

	var data []byte
	data, err = syscall.Mmap(fd, 0, 512, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalln("mmap:", err)
	}

	log.Print(string(data))
}
