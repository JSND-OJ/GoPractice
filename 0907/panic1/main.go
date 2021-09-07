package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0알 수 없습니다.")
	}
	fmt.Printf("%d / %d = %d", a, b, a/b)
}

func main() {
	divide(9, 0)
	divide(9, 3)
}
