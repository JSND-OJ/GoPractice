package main

import "fmt"

func main() {
	dan := 2
	b := 1
	for {
		for{
			fmt.Println("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 {
				break
			}
		}
		b = 1
		dan++ 
			break
		}
	}
	fmt.Println("for문이 종료되었습니다")
}