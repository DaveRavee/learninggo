package main

import "fmt"

type person struct {
	first string
	last string
	favIceCream []string
}


func main() {

	p1 := person{
		first: "Dave",
		last: "Jeff",
		favIceCream: []string{"mint", "vanilla", "cookie"},
	}

	p2 := person{
		first: "Lel",
		last: "kek",
		favIceCream: []string{"cheese", "logan", "mango"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first)
	fmt.Println(p1.last)

	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)

	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}



}
