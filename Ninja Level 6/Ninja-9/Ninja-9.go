package main

import "fmt"

func main() {

	a := func(b []int) int {
		if len(b) == 0 {
			return 1
		}
		if len(b) == 1 {
			return b[0]
		}
		return b[0] + b[len(b)-1]
	}

	c := foo(a, []int{22, 23, 25, 23, 25, 11})
	fmt.Println(c)
}

func foo(x func(y []int) int, z []int) int {
	i := x(z)
	i += 12
	return i
}


//func foo(x func()) {
//	x()
//}
//
//func bar() func() {
//	return func() {
//		fmt.Println("hello")
//	}
//}