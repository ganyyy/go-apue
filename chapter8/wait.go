// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
)

func main() {

	var wStatus = func(status *syscall.WaitStatus) string {
		var ss []string
		if status.Exited() {
			ss = append(ss, fmt.Sprintf("is exit:%v", status.ExitStatus()))
		}
		if status.Signaled() {
			ss = append(ss, status.Signal().String())
		}

		if status.Stopped() {
			ss = append(ss, status.StopSignal().String())
		}

		if status.CoreDump() {
			ss = append(ss, "core dump")
		}

		return strings.Join(ss, ",")
	}

	{
		var pid, _, _ = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if pid == 0 {
			time.Sleep(time.Second)
			os.Exit(10)
		}

		var status syscall.WaitStatus
		var wpid, _ = syscall.Wait4(int(pid), &status, 0, nil)
		fmt.Println(pid, wpid, wStatus(&status))
	}

	{
		var pid, _, _ = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if pid == 0 {
			time.Sleep(time.Second)
			panic("")
		}

		var status syscall.WaitStatus
		var wpid, _ = syscall.Wait4(int(pid), &status, 0, nil)
		fmt.Println(pid, wpid, wStatus(&status))
	}

	{
		var pid, _, _ = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if pid == 0 {
			time.Sleep(time.Second)
			var a, b int
			a /= b
		}

		var status syscall.WaitStatus
		var wpid, _ = syscall.Wait4(int(pid), &status, 0, nil)
		fmt.Println(pid, wpid, wStatus(&status))

	}
}
