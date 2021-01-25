package main

import "fmt"

func main() {

	a := factorial(4)
	fmt.Println(a)
}

func factorial(x int)  int{
	y := 1
	for ; x > 0; x-- {
		y *= x
	}
	return y
}