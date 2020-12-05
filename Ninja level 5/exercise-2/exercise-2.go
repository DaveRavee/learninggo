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

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for _, v := range m {

		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
	}

}
