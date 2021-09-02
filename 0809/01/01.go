package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct { //은행계좌
	balance int //잔액 (메모리 보호대상) Mutex
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) { //인출
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) { //입금
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

var accounts []*Account //slice 배열 (메모리 보호대상) Mutex

func Transfer(sender, receiver int, money int) {
	accounts[sender].mutex.Lock()
	// time.Sleep(1000 * time.Millisecond)  //dead lock()발생
	accounts[receiver].mutex.Lock()
	// fmt.Println("Lock", receiver)   //어떤 락을 잡았는지...출력을 해본다.데드락현상 확인
}

func GetTotalBalance() int { //전체 잔액량
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func RandomTransfer() { //랜덤함수 레퍼런스
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 { //샌더 잔액여부 확인
			break //샌더 잔액이 있으면 멈춘다
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ { //20개 어카운트 배열 1000 지정
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	go func() {
		for {
			Transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			Transfer(1, 0, 100)
		}
	}()

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
