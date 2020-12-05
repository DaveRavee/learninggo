package main

import "fmt"

func main() {
	x := []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	a := x[0:5]
	b := x[5:10]
	c := x[2:7]
	d := x[1:6]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
