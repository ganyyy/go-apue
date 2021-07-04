// +build ignore

package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
}

func (m *MyStruct) String() string {
	return reflect.TypeOf(m).Elem().Name()
}

func main() {
	fmt.Println(&MyStruct{})
	var s = -1
	var m = make(map[int]bool, s)

	fmt.Println(m[0])
}
