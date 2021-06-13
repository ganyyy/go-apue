// +build ignore

package main

import (
	"fmt"
	"go-apue/helper"
	"os"
	"syscall"
)

var NTotal int

type Count int

func (c Count) Percent(name string) {
	fmt.Printf("%s %7d, %5.2f", name, c, float32(c)/float32(NTotal))
}

func main() {
	if len(os.Args) != 2 {
		helper.PanicIfError("exit", fmt.Errorf("args num error"))
	}

	var (
		NReg   Count
		NDir   Count
		NBlk   Count
		NChr   Count
		NFifo  Count
		NSlink Count
		NSock  Count
		NTot   Count
	)

	defer func() {
		NTotal = int(
			NReg +
				NDir +
				NBlk +
				NChr +
				NFifo +
				NSlink +
				NSock +
				NTot)
		if NTotal == 0 {
			NTotal = 1
		}
		NReg.Percent("NReg")
		NDir.Percent("NDir")
		NBlk.Percent("NBlk")
		NChr.Percent("NChr")
		NFifo.Percent("NFifo")
		NSlink.Percent("NSlink")
		NSock.Percent("NSock")
		NTot.Percent("NTot")
	}()

	var path = os.Args[1]
	var err error
	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		helper.PanicIfError("Stat", err)
	}

	if stat.Mode&syscall.S_IFDIR == 0 {
		helper.PanicIfError("type error", fmt.Errorf("%v not dir", path))
	}

	//TODO 暂时不想写了..
}
