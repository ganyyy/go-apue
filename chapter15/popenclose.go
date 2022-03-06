// +build ignore

package main

import (
	"os"
	"os/exec"
)

func main() {
	var cmd = exec.Command("./myuclc")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Start()

	cmd.Wait()
}
