package main

import "fmt"

func main() {
	x := []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
