package main

import (
	"fmt"
)

type Dupa struct{}

func (d *Dupa) dupa(msg string) *Dupa {
	fmt.Println(msg)
	return d
}

func main() {
	d := Dupa{}
	defer d.dupa("2").dupa("3")
	d.dupa("1")
}
