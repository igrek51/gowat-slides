package main

import (
	"fmt"
)

type gopher struct {
	Num int
}

func (d *gopher) print() {
	fmt.Println(d)
}

func main() {
	gophers := []gopher{gopher{1}, gopher{4}, gopher{6}, gopher{9}}

	for _, g := range gophers {
		defer g.print()
	}
}
