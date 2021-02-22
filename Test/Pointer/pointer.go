package main

import "fmt"

func main() {
	x := 23

	fmt.Println("The value of x is", x)
	fmt.Println("The memory location of x is at", &x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	y := &x
	*y = 42

	fmt.Println("Using pointers to change value of x, the value is now", x)

}
