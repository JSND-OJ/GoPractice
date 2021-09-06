package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	m := make(map[int]Student)

	m[5] = Student{"aaa", 43}
	m[7] = Student{"bbb", 76}
	m[18] = Student{"ccc", 99}

	fmt.Print("번", m[5])
}
