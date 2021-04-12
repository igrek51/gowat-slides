package main

import (
	"fmt"
)

func main() {
	name := "dupa"
	switch name {
	case "dupa":
	case "Dupa":
		fmt.Println("hello dupa.")
		return
	default:
		fmt.Println("default")
		return
	}
	fmt.Println("idź pan w chuj.")
}
