package main

import "fmt"

func main() {
	var age = 60

	if age >= 10 && age <= 15 {
		fmt.Println("로리")
	} else if age > 30 || age < 20 {
		fmt.Println("20대가 아니네여")
	} else {
		fmt.Println("좋네")
	}
}
