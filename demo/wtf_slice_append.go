package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 3)
	b := a
	fmt.Println("a =", a, len(a), cap(a))
	fmt.Println("b =", b, len(b), cap(b))

	a = append(a, 1)
	fmt.Println("---")
	fmt.Println("a =", a, len(a), cap(a))
	fmt.Println("b =", b, len(b), cap(b))

	b = append(b, 2)
	fmt.Println("---")
	fmt.Println("a =", a, len(a), cap(a))
	fmt.Println("b =", b, len(b), cap(b))
}
