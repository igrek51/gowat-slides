package main

import (
	"fmt"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("type:%T; value:%v; value#:%#v", v, v, v)
}

func Newstr(s string) *string {
	return &s
}

func main() {
	l := make([]*string, 8)

	fmt.Println("l =", typeof(l))

	l[1] = Newstr("du")
	l[3] = Newstr("pa")

	fmt.Println("l =", typeof(l))
}
