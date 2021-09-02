package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		for j := 0; j < i + 1; j++ {
			fmt.Println("*")
		}
		fmt.Println()
	}
	for i := 0; i < 2; i-- {
		for j := 0; j < 2 - i: j++ {
			fmt.Println("*")
		}
	}
}
