package main

import (
	"fmt"
)

type Clazz struct {
	field1 string
	field2 string
}

func main() {
	l := []string{
		"Trailing",
		"comma",
		"is cool",
	}
	_ = Clazz{
		field1: "tabulated",
		field2: "tabulated",
	}

	if l != nil {
		fmt.Println("no parentheses needed")
	}
}
