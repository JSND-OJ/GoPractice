package main

import "fmt"

const M = 10

func hash(d int) int {
	return d % M
}

func main() {

	//배열사용
	a := [M]int{}

	a[hash(0)] = 10
	a[hash(1)] = 11
	a[hash(2)] = 12
	a[hash(3)] = 13
	a[hash(4)] = 14
	a[hash(5)] = 15
	a[hash(6)] = 16
	a[hash(7)] = 17
	a[hash(8)] = 18
	a[hash(9)] = 19

	//delete(a, hash(0)) 안먹힘
	//ac, ok := a[hash(1)] 이것도 안먹힘
	//fmt.Println(ac, ok)

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Println()

	//맵사용
	m := make(map[int]int)

	m[hash(0)] = 10
	m[hash(1)] = 11
	m[hash(2)] = 12
	m[hash(3)] = 13
	m[hash(4)] = 14
	m[hash(5)] = 15
	m[hash(6)] = 16
	m[hash(7)] = 17
	m[hash(8)] = 18
	m[hash(9)] = 19

	delete(m, hash(0))

	mc, ok := m[hash(1)]
	fmt.Println(mc, ok)

	for k, v2 := range m {
		fmt.Println(k, v2)
	}
}
