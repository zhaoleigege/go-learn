package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls")
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprint(os.Stderr, string(result))
	}

	fmt.Println(string(result))
}
