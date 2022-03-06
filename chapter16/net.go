// +build ignore

package main

import (
	"log"
	"net"
)

func main() {

	log.Println(net.LookupPort("tcp", "ftp"))
	log.Println(net.LookupCNAME("www.baidu.com"))
	log.Println(net.LookupHost("localhost"))

}
