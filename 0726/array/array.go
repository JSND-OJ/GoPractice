package main

import "fmt"

func main() {
	s := "잠온당 sibal"

	s2 := []rune(s)
	fmt.Println("len(s2)", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ", ")
	}
}
