package main

import (
	"fmt"
)

type Container struct {
	m map[string]string
}

func main() {
	a := map[string]string{}
	a["a"] = "A"

	b := a
	b["b"] = "B"

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := Container{
		m: map[string]string{},
	}
	c.m["c"] = "C"

	d := c
	d.m["d"] = "D"

	fmt.Println("c =", c)
	fmt.Println("d =", d)

}
