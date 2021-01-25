package main

import "fmt"

func main() {

	func() {
		fmt.Println("my nuts")
	}()

	func(x int) {
		fmt.Println(x)
	}(22)
}

