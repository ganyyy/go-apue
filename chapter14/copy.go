// +build ignore

package main

import (
	"log"
	"os"
	"syscall"
)

func main() {

	const (
		BlockSize = 1024 * 1024 * 4 // 单次拷贝4M的数据
	)

	var fin, fout int
	var src, dst []byte
	var copySize int64
	var stat syscall.Stat_t
	var fsz int64
	var err error

	if len(os.Args) != 3 {
		log.Printf("please input src file and dst file path")
		os.Exit(1)
	}

	fin, err = syscall.Open(os.Args[1], syscall.O_RDONLY, 0)
	if err != nil {
		log.Panicf("open %s error %v", os.Args[1], err)
	}

	fout, err = syscall.Open(os.Args[2], syscall.O_RDWR|syscall.O_CREAT|syscall.O_TRUNC, uint32(os.ModePerm))
	if err != nil {
		log.Panicf("open %s error %v", os.Args[2], err)
	}

	err = syscall.Fstat(fin, &stat)
	if err != nil {
		log.Panicf("state fin %s error %v", os.Args[1], err)
	}

	err = syscall.Truncate(os.Args[2], stat.Size)
	if err != nil {
		log.Panicf("truncate %s error %v", os.Args[2], err)
	}

	for fsz < stat.Size {
		if remain := stat.Size - fsz; remain > BlockSize {
			copySize = BlockSize
		} else {
			copySize = remain
		}

		src, err = syscall.Mmap(fin, fsz, int(copySize), syscall.PROT_READ, syscall.MAP_SHARED)
		if err != nil {
			log.Panicf("mmap from fin error:%v", err)
		}

		dst, err = syscall.Mmap(fout, fsz, int(copySize), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
		if err != nil {
			log.Panicf("mmap to fout error:%v", err)
		}

		copy(dst, src)
		syscall.Munmap(src)
		syscall.Munmap(dst)
		fsz += copySize
	}

	os.Exit(0)

}
