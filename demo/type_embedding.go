package main

import "fmt"

type Rick struct {
	morty string
}

func (r Rick) whoAmI() {
	fmt.Println("Rick Sanchez")
}

type PickleRick struct {
	Rick
	pickles int
}

func (r PickleRick) whoAmI() {
	fmt.Println("I'm a pickle!")
}

func main() {
	rick := PickleRick{Rick: Rick{morty: "evil Morty"}, pickles: 1}
	fmt.Println(rick)
	fmt.Println(rick.morty, rick.Rick.morty)
	rick.whoAmI()
	rick.Rick.whoAmI()
}
