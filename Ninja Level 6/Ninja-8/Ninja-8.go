package main

import "fmt"

func main() {
	x := nuts()
	fmt.Println(x())
}



func nuts() func() int {
	return func() int {
		return 123
	}
}