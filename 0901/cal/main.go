package main

import "fmt"

//숫자입력
func inputNumber() (int, int) {
	x, y := 0, 0

	for {
		fmt.Println("숫자입력")

		_, err := fmt.Scanf("%d,%d\n", &x, &y)

		if err != nil {
			fmt.Println("다시입력")
		} else {
			break
		}
	}

	return x, y

}

//연산소자입력
func inputOperator() string {
	op := ""
	for {
		fmt.Println("연산자입력")

		_, err := fmt.Scanf("%s \n", &op)
		if err != nil {
			fmt.Println("다시입력")
		} else {
			break
		}
	}
	return op
}

//숫자와 연산소자로 연산
func calculate(x, y int, op string) int {
	result := 0

	switch op {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	}
	return result
}

func main() {
	x, y := inputNumber()
	op := inputOperator()
	result := calculate(x, y, op)

	fmt.Printf("%d %s %d = %d", x, op, y, result)
}
