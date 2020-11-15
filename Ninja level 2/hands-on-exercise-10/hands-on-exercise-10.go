package main

import "fmt"

const (
	thisyear = 2020 + iota
	nextyear = 2020 + iota
	nextnextyear = 2020 + iota
	nextnextnextyear = 2020 + iota
	nextnextnextnextyear = 2020 + iota
)

func main() {
	fmt.Println(thisyear)
	fmt.Println(nextyear)
	fmt.Println(nextnextyear)
	fmt.Println(nextnextnextyear)
	fmt.Println(nextnextnextnextyear)

}
