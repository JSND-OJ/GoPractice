package main

import "fmt"

func main() {
	a := 3
	b := 5

	//사칙연산
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b + %d\n", a-b)
	fmt.Printf("a * b + %d\n", a*b)
	fmt.Printf("a / b + %d\n", a/b)

	//사칙연산
	fmt.Printf("a & b = %d\n", a&b)
	fmt.Printf("a | b = %d\n", a|b)
	fmt.Printf("a ^ b = %d\n", a^b)
	fmt.Printf("a &^ b = %d\n", a&^b)

	//쉬프트연산
	fmt.Printf("a >> b = %d\n", a>>b)
	fmt.Printf("a << b = %d\n", a<<b)

	//논리연산
	fmt.Println("a > b = ", a > b)
	fmt.Println("a < b = ", a < b)
	fmt.Println("a >= b = ", a >= b)
	fmt.Println("a <= b = ", a <= b)
	fmt.Println("a == b = ", a == b)
	fmt.Println("a != b = ", a != b)

	//나머지 연산
	fmt.Print("a % b = ", a%b)
}
