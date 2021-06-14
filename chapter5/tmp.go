// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var name, _ = ioutil.TempDir("", "123*")

	fmt.Printf(name)

	var t, _ = ioutil.TempFile("", "a*")
	t.WriteString("123456")
	t.Close()

}
