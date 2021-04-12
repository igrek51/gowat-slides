package main

import (
	"fmt"
)

type alias32 int32

func main() {
	var aliased alias32 = 5

	whatis(int32(5))
	whatis(aliased)
}

func whatis(value interface{}) {
	switch value := value.(type) {
	case int32:
		fmt.Println(value, "int32")
	default:
		fmt.Println(value, "idz pan w chuj")
	}
}
