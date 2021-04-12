package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 1)

	go sendDupa(messages, 1)
	go sendDupa(messages, 2)
	go sendDupa(messages, 3)

	time.Sleep(1 * time.Second)
	fmt.Println("len:", len(messages), "capacity:", cap(messages))

	for i := 0; i < 3; i++ {
		fmt.Println("received:", <-messages)
	}
}

func sendDupa(messages chan string, id int) {
	fmt.Println("sending...")
	messages <- fmt.Sprintf("dupa %d", id)
	fmt.Println("sent")
}
