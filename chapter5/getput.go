// +build ignore

package main

import (
	"bufio"
	"os"
)

func main() {
	var buf = bufio.NewReader(os.Stdin)

	for line, err := buf.ReadSlice('\n'); err == nil; line, err = buf.ReadSlice('\n') {
		os.Stdout.Write(append([]byte("123:"), line...))
	}

}
