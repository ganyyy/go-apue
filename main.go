// +build ignore

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type MyStruct struct {
}

func (m *MyStruct) String() string {
	return reflect.TypeOf(m).Elem().Name()
}

func main() {
	fmt.Println((*MyStruct)(unsafe.Pointer(nil)).String())

	fmt.Println(reflect.TypeOf(nil).Elem().Name())
}
