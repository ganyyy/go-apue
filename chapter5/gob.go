// +build ignore

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type TestVal struct {
	Name string
}

func main() {

	var buf = bytes.NewBuffer(nil)

	var encoder = gob.NewEncoder(buf)

	encoder.Encode(TestVal{
		Name: "12345",
	})

	var t = TestVal{}
	var decoder = gob.NewDecoder(buf)
	decoder.Decode(&t)

	fmt.Println(t, buf.Len())

	var f, _ = os.Open("123")
	f.Fd()
}
