package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num01 int
	var result int
	var no1 int
	var no2 int
	var score int
	rand.Seed(time.Now().UnixNano())
	score = 0
	fmt.Println("")
	time.Sleep(1000)
	for {
		fmt.Println("")
		time.Sleep(1000)
		fmt.Println("")
		no1 = rand.Intn(10)
		no2 = rand.Intn(10)
		fmt.Println("")
		time.Sleep(1000)
		for {
			fmt.Println("")
			_, err := fmt.Scanf("%d\n", no1, no2)

			if err != nil {
				fmt.Println("")
				continue
			} else {
				break
			}
		}
		if num01 == result {
			fmt.Println("")
			score++
		} else {
			fmt.Println("")
			score--
		}
		if score == 5 {
			fmt.Println("")
			break
		}
	}
}
