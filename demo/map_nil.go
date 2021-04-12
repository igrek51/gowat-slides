package main

import (
	"fmt"
)

func main() {
	var m map[string]string
	fmt.Printf("m['dupa'] = '%s'\n", m["dupa"])
	m["dupa"] = "dupa"
}
