package main

import "fmt"

func add(x, y string) (string, string) {
	return x, y
}

func main() {
	a, b := add("hello", "world")
	fmt.Println(a, b)
}
