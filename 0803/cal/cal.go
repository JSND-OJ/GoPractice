package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("첫 번째 숫자")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	num01, _ := strconv.Atoi(line)

	fmt.Println("첫 번째 숫자")
	reader01 := bufio.NewReader(os.Stdin)
	line01, _ := reader.ReadString('\n')
	line01 = strings.TrimSpace(line01)
	num02, _ := strconv.Atoi(line01)

	fmt.Println("연산자")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)
	result := 0
	// if op == "+" {
	// 	result = num01 + num02
	// 	fmt.Printf("%v + %v = %v", num01, num02, result)
	// } else if op == "-" {
	// 	result = num01 - num02
	// 	fmt.Printf("%v - %v = %v", num01, num02, result)
	// } else if op == "*" {
	// 	result = num01 * num02
	// 	fmt.Printf("%v * %v = %v", num01, num02, result)
	// } else if op == "/" {
	// 	result = num01 / num02
	// 	fmt.Printf("%v / %v = %v", num01, num02, result)
	// } else {
	// 	fmt.Println("연산자가 아닙니다")
	// }
	switch op {
	case "+": 
	result = num01 + num02
	
	case "-":
		result = num01 - num02
	
	case "*":
		result = num01 * num02
	
	case "?":
		result = num01 / num02
	


		}
	}
}
