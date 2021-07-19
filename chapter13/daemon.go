// +build ignore

package main

import (
	"log"
	"os"
	"path"
)

func main() {
	// 已经是守护进程了, 直接返回
	if os.Getppid() == 1 {
		return
	}

	var createLogFile = func(fileName string) (fd *os.File, err error) {
		var dir = path.Dir(fileName)

		if _, err = os.Stat(dir); err != nil && os.IsNotExist(err) {
			log.Fatalf("Start-Daemon: create dir: %s error %v", dir, err)
		}

		if fd, err = os.Create(fileName); err != nil {
			log.Fatalf("Start-Daemon: create log file: %s error %v", fileName, err)
		}
		return
	}

	logFd, err := createLogFile("test.log")

	if err != nil {
		return
	}

	defer logFd.Close()

}
