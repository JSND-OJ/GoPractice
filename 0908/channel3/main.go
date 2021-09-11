package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //2.데이터를 계속 기다림 ->for문이 끝나서 더이상 채널이 들어오지 않지만 그래도 계속 기다림.
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() // . 실행되지 않음
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2 // 1.데이터를 넣음
	}
	close(ch) //3.채널을 닫아준다.
	//wg.Wait() // 작업 완료를 기다림. 근데 필요없지않음?
}
