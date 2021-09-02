package main

import "fmt"

func main() {

	a, b := fun1(2, 3)

	fmt.Println(a, b)

}

func fun1(x, y int) (int, int) {
	fun2(x, y)
	return y, x
}

func fun2(x, y int) {
	fmt.Println(("func2", x, y))
}
