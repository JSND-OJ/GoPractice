package main

import "fmt"

func main() {

	for i := 2; i < 10; i++ {
		if (i % 2) == 0 {
			fmt.Println(i, "단")

			for j := 1; j < 10; j++ {
				if (j % 3) == 0 {
					fmt.Printf("%d * %d = %d\n", i, j, i*j)
				}
			}
		}
	}
}
