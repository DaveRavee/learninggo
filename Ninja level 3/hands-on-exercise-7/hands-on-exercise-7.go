package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("The number %v, is a factor of two\n", i)
		}
		if i%3 == 0 {
			fmt.Printf("The number %v, is a factor of three\n", i)
		}
		if i%4 == 0 {
			fmt.Printf("The number %v, is a factor of four\n", i)
		}
	}
}
