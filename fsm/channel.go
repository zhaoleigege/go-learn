package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func worker(jobChan <-chan string) {
	for job := range jobChan {
		process(job)
	}
}

func process(str string) {
	time.Sleep(3 * time.Second)
	fmt.Println("value: ", str)
}

func main() {
	jobChan := make(chan string, 100)
	workCount := 5

	for i := 0; i < workCount; i++ {
		go worker(jobChan)
	}

	for i := 0; i < 200; i++ {
		jobChan <- fmt.Sprintf("test-%d", i)
	}

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		jobChan <- input.Text()
	}
}
