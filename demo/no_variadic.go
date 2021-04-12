package main

func variadic(elems ...interface{}) int {
	return 0
}

func main() {
	slice := []string{"a", "b"}
	variadic("a", "b")
	variadic(slice...) // compilation error
	// you can do: variadic(slice) - but it's not we meant
}
