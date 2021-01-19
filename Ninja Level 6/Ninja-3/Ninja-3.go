package main

import "fmt"

func main() {
	defer bar()
	foo()
}

func foo() {
	fmt.Println("foo has ran")
}

func bar() {
	fmt.Println("bar has ran")
}