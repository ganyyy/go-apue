// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, ","))
	fmt.Println(strings.Join(os.Environ(), ","))
}
