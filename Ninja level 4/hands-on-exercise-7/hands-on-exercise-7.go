package main

import "fmt"

func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helllo, James"}

	a := [][]string{x, y}

	for i, v := range a {
		fmt.Println(i, v)
	}


}
