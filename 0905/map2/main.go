package main

import "fmt"

type Product struct {
	Name  string
	Price float32
}

func main() {
	m := make(map[int]Product)

	m[14] = Product{"aaa", 1000}
	m[321] = Product{"bbb", 999.9}
	m[134] = Product{"ccc", 99.9}
	m[11234] = Product{"ddd", 24.49}
	m[1431414] = Product{"eee", 1999.99}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
