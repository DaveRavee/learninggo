package main

import "fmt"

func main() {
	i := increment()
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())
	fmt.Println(i())

}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
