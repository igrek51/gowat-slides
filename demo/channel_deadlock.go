package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	ch <- "dupa"

	fmt.Println("end")
}
