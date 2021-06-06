// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var s string
	var err error

	for _, err = fmt.Scanf("%s", &s); err == nil; _, err = fmt.Scanf("%s", &s) {
		var cmd = exec.Command(s)
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			log.Printf("exec %v error %v", s, err)
		}
		_ = cmd.Wait()
	}

	log.Println("shell exit")
	log.Println("shell exit")
	log.Println("shell exit")
}
