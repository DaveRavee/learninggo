package main

import "fmt"

func main() {

	numb := []int{1, 2, 3, 4, 5}
	s1 := foo(numb...)
	fmt.Println(s1)

	s2 := bar(numb)
	fmt.Println(s2)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(y []int) int {
	total := 0
	for _, v := range y {
		total += v
	}
	return total
}