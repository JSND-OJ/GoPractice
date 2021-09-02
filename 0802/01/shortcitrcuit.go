package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.println("IncreaseAndReturn()", cnt) {
	cnt ++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1증가")
	}
	if true || IncreaseAndReturn() <5 {
		fmt.Println("2증가")
	}

fmt.Println("cnt:", cnt)
}