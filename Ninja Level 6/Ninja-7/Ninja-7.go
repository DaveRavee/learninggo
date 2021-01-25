package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := func() {
		fmt.Println("Talk some smack")
	}
	y()
}

func foo() int {
	return 4
}