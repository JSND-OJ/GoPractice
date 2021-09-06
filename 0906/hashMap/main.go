package main

import "fmt"

const M = 10

func hash(d int) int {
	return d % M
}

func main() {

	//배열사용
	a := [M]int{}

	a[hash(23)] = 10
	fmt.Printf("%d = %d\n", 23, a[hash(23)])
	a[hash(33)] = 11
	a[hash(259)] = 50

	fmt.Printf("%d = %d\n", 23, a[hash(23)])
	fmt.Printf("%d = %d\n", 33, a[hash(33)])
	fmt.Printf("%d = %d\n", 259, a[hash(259)])
	fmt.Println()

	//맵사용
	m := make(map[int]int)

	m[hash(23)] = 10
	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	m[hash(33)] = 11
	m[hash(259)] = 50

	fmt.Printf("%d = %d\n", 23, m[hash(23)])
	fmt.Printf("%d = %d\n", 33, m[hash(33)])
	fmt.Printf("%d = %d\n", 259, m[hash(259)])
}
