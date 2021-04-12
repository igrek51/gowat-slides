package main

import "fmt"

// type alias
type bignumber int64

// struct
type Empty struct{}
type Address struct{ host string }
type Address2 struct {
	host string
}

// interface
type I interface{}
type I2 interface {
	match(input string) bool
}

// func
type Matcher func(input string) bool

func printIt(it struct{ host string }) {
	fmt.Println(it, it.host)
}

func main() {
	printIt(Address{host: "localhost"})
	printIt(Address2{host: "localhost"})

	var aliased bignumber = 5_000_000_000
	fmt.Println(aliased)
	_ = []Matcher{}
}
