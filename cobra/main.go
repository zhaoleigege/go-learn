package main

import "github.com/buse/cmd"

// 参考资料 https://cobra.dev/
func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
