package main

import "fmt"

type Dupa struct{}

func dupaBuilder() *Dupa {
	return &Dupa{}
}

func printProvided(provider func() interface{}) {
	fmt.Print(provider())
}

func main() {
	printProvided(dupaBuilder)
}
