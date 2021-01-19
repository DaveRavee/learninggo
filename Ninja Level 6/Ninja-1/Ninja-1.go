package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y, z := bar()
	fmt.Println(y, z)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 6, "hello"
}