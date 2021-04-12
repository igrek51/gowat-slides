package main

import (
	"fmt"
)

type Dupa struct{}

func (d *Dupa) print() {
	fmt.Println("dupa")
}

func NewDupa() *Dupa {
	return nil
}

func main() {
	NewDupa().print()
}
