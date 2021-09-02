package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}

	slice2 := append(slice1, 4, 5)
	//cap함수르 이용해 슬라이스 capacity갑을 알아볼 수 있습니다.
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("After Change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("After Change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice1), cap(slice2))

}
