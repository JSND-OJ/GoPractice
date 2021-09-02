package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
	fmt.Println(array2)
}
func changeSlace(slice2 []int) {
func changeSlace(slice2 []int) {
	slice2[2] = 200
	fmt.Println(slice2)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlace(slice)

	fmt.Println("array:", array)
	fmt.Println("slice:", slice)
}

