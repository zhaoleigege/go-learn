package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func login() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("数据: %d", v)
		default:
			//time.Sleep(time.Second)
		}
	}
}

func main() {
	var isCpu bool
	var isMem bool
	flag.BoolVar(&isCpu, "cpu", false, "cpu检测")
	flag.BoolVar(&isMem, "mem", false, "memory检测")
	flag.Parse()

	if isCpu {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			panic(err)
		}

		if err = pprof.StartCPUProfile(file); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 10; i++ {
		go login()
	}

	time.Sleep(20 * time.Second)

	if isMem {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			panic(err)
		}

		if err := pprof.WriteHeapProfile(file); err != nil {
			panic(err)
		}
		file.Close()
	}
}
