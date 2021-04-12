package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	go func() {
		panic("dupa")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end")
}
