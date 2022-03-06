// +build ignore

package main

import "fmt"

func main() {
	var c byte
	for n, err := fmt.Scanf("%c", &c); n != 0 && err == nil; n, err = fmt.Scanf("%c", &c) {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		} else if c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		fmt.Printf("%c", c)
	}
}
