package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // 채널 생성

	wg.Add(1)
	go square(&wg, ch) // 고루틴 생성
	ch <- 9            // 채널에 데이터 넣음
	wg.Wait()          // 작업이 완료되기를 기다림
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch // 데이터 넣음

	time.Sleep(time.Second) // 1초 대기
	fmt.Printf("Square %d\n", n*n)
	wg.Done()
}
