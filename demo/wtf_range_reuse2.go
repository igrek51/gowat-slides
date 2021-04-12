package main

import (
	"fmt"
	"time"
)

func main() {
	values := []string{"dupa1", "dupa2", "dupa3"}

	for _, val := range values {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(val)
		}()
	}

	time.Sleep(2 * time.Second)
}

func fixed_by_goroutine_param(values []string) {
	for _, val := range values {
		go func(val string) {
			time.Sleep(1 * time.Second)
			fmt.Println(val)
		}(val)
	}
}

func fixed_by_reassign(values []string) {
	for _, val := range values {
		reassigned := val
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println(reassigned)
		}()
	}
}
