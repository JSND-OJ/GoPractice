package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func PrintNumbers() {
	for i := 1; 1 < 6; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
}
