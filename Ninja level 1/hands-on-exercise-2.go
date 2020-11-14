package main

import "fmt"

//compiler assigns 'zero values' to these values that do not have values stored
var x int
var y string
var z bool

func main() {

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
