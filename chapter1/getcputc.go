// +build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// var inputBuf = bufio.NewReader(os.Stdin)
	var outputBuf = bufio.NewWriter(os.Stdout)

	var input string
	var err error

	for _, err = fmt.Scan(&input); err == nil; _, err = fmt.Scan(&input) {
		outputBuf.WriteString(input)
		outputBuf.WriteByte('\n')
		outputBuf.Flush()
	}

	log.Printf("read buf error:%v", err)
}
