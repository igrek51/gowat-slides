package main

import (
	"fmt"
	"reflect"
)

type I interface{}

type Dupa struct{}

func NewDupa() *Dupa {
	return nil
}

func NewI() I {
	return NewDupa()
}

func getValue() interface{} {
	var a *string = nil
	return a
}

func main() {
	// short example
	fmt.Println(getValue() == nil)
	fmt.Println(getValue().(*string) == nil)

	fmt.Println("*Dupa == nil:", NewDupa() == nil)

	dupa := NewI()
	fmt.Println("dupa == nil:", dupa == nil)
	fmt.Println("typeof(dupa):", typeof(dupa))
	fmt.Println("dupa.(I) == nil:", dupa.(I) == nil)
	fmt.Println("dupa.(interface{}) == nil:", dupa.(interface{}) == nil)
	fmt.Println("dupa.(*Dupa) == nil:", dupa.(*Dupa) == nil)
	fmt.Println("reflect.ValueOf(dupa).IsNil():", reflect.ValueOf(dupa).IsNil())
}

func typeof(v interface{}) string {
	return fmt.Sprintf("type:%T; value:%v; %#v", v, v, v)
}
