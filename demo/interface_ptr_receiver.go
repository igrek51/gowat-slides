package main

import (
	"fmt"
)

type I interface {
	print()
}

type Dupa1 struct{}
type Dupa2 struct{}

func (d *Dupa1) print() {
	fmt.Println("*Dupa1")
}

func (d Dupa2) print() {
	fmt.Println("Dupa2")
}

func NewDupa2() I {
	return Dupa2{}
}
func NewDupa1Ptr() I {
	return &Dupa1{}
}
func NewDupa2Ptr() I {
	return &Dupa2{}
}

func main() {
	NewDupa2().print()
	NewDupa1Ptr().print()
	NewDupa2Ptr().print()
}
