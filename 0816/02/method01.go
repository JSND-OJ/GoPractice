package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100} //a의 타입은 *account 타입

	fmt.Printf("%d\n", a.balance)

	withdrawFunc(a, 30)

	fmt.Printf("%d\n", a.balance)

	a.withdrawMethod(30)

	fmt.Printf("%d \n", a.balance)
}
