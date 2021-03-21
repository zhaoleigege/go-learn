package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("pprof-file.pb.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	mf, err := os.Create("memory.pb.gz")
	if err != nil {
		panic(err)
	}
	pprof.WriteHeapProfile(mf)
	defer mf.Close()

	fmt.Println(calc())
}

func calc() int {
	sum := 0
	for i := 0; i < 1000*1000; i++ {
		arr := make([]int, i)
		sum = sum * (i + 1) / 2
		if len(arr) > 1000 {

		}
	}

	return sum
}
