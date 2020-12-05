package main

import "fmt"

func main() {
	bank := struct {
		accountNumb int
		first string
		last string
		balance int
	}{
		accountNumb: 1234,
		first: "Dave",
		last: "Jeff",
		balance: 22,
	}

	fmt.Println(bank.accountNumb)
	fmt.Println(bank.first, bank.last)
	fmt.Println(bank.balance)


}
