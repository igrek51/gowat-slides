package main

import (
	"fmt"
	"time"
)

func somethingVeryLong(input int) int {
	time.Sleep(1 * time.Second)
	return input * 2
}

func doItAsync(input int) <-chan int {
	resultCh := make(chan int)
	go func() {
		defer close(resultCh)
		resultCh <- somethingVeryLong(input)
	}()
	return resultCh
}

func main() {
	job1 := doItAsync(1)
	job2 := doItAsync(10)
	job3 := doItAsync(21)

	fmt.Println(<-job1)
	fmt.Println(<-job2)
	fmt.Println(<-job3)
}
