package main

import "fmt"

func main() {
	go doNothing()
	select {}
}

func doNothing() {
	events := make(chan string)
	for {
		select {
		case <-events:
			fmt.Println("event received")
		default:
		}
	}
}
