package main

import "fmt"

func main() {
	// a := 3

	// if a > 3 {
	// 	fmt.Println("정답")
	// } else if a < 2 {
	// 	fmt.Println("노답")
	// } else {
	// 	fmt.Println("정답")
	// }

	// switch a {
	// case 1:
	// 	fmt.Println("정답")
	// case 2:
	// 	fmt.Println("정답")
	// case 3:
	// 	fmt.Println("정답")
	// default:
	// 	fmt.Println("정답")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d\t", i)
	// }

	//*표 찍기 문제

	//구구단

	for i := 2; i < 10; i++ {
		fmt.Println(i, "단")
		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
	}

}
