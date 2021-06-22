// +build ignore

package main

import (
	"fmt"
	"os"
)

func main() {
	var f, err = os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
		f.WriteString("tttt")
	}
}
