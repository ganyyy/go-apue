package helper

import (
	"fmt"
	"log"
	"syscall"
)

const PATH = "/tmp/unix_file_transfer"

//SendFD 向指定的目的unix套接字fd 发送 sendto 文件描述符
func SendFD(fd int, sendto int, msg string) error {
	return syscall.Sendmsg(fd, []byte(msg), syscall.UnixRights(sendto), nil, 0)
}

//RecvFD 从指定的unix套接字fd上
func RecvFD(fd int) (int, string, error) {
	var oob [32]byte
	var msg [128]byte
	n, oobn, _, _, err := syscall.Recvmsg(fd, msg[:], oob[:], 0)
	if err != nil {
		return -1, "", err
	}
	scms, err := syscall.ParseSocketControlMessage(oob[:oobn])
	if err != nil {
		return -1, "", err
	}
	fds, err := syscall.ParseUnixRights(&(scms[0]))
	if err != nil {
		return -1, "", err
	}
	log.Printf("recv fds:%v", fds)
	if len(fds) == 0 {
		return -1, "", fmt.Errorf("not have fds")
	}
	return fds[0], string(msg[:n]), nil
}
