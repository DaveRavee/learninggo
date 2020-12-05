package main

import "fmt"

func main() {
	arr := [5]int {1, 2, 3, 4, 5}
	//a := arr[0]
	//b := arr[1]
	//c := arr[2]
	//d := arr[3]
	//e := arr[4]

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", arr)
}
