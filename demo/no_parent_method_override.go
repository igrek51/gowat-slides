package main

import "fmt"

func main() {
	var s son = son{}
	s.whoami()
	s.askYourself()
}

type father struct {
}

func (self father) whoami() {
	fmt.Println("I'm your father")
}
func (self father) askYourself() {
	fmt.Println("ask yourself...")
	self.whoami()
}

type son struct {
	father
}

func (self son) whoami() {
	fmt.Println("son")
}
