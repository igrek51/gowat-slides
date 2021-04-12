package main

type Base interface {
}

type Extended interface {
	Base
}

func main() {
	bases := []Base{nil}
	extends := []Extended{nil}

	bases[0] = extends[0] // single elements are assignable
	bases = extends       // ...but slices are not - type mismatch
}
