package main

import "fmt"

func main() {
	var temp int
	var rain int
	fmt.Scanf(&temp, &rain)

	if temp > 25 {
		if rain >= 80 {
			fmt.Println("덥고비가옵니다")
		} else if rain >= 20 {
			fmt.Println("덥고습합니다")
		} else {
			fmt.Println("야외활동하기좋습니다")
		}
	} else if temp < 10 || rain >= 80 {
		fmt.Println("야외활동하기좋지않습니다")
		fmt.Println("좋은날씨입니다")
	} else {
		fmt.Println("좋은날씨입니다")
	}
}
