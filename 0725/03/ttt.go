package main

import (
 "fmt"
 "math/rand"
 "sync"
 "time"
)

type Account struct {
 balance int
 mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
 a.mutex.Lock()
 a.balance -= val
 a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
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

var accounts []*Account
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
 globalLock.Lock()
 accounts[sender].Widthdraw(money)
 accounts[receiver].Deposit(money)
 globalLock.Unlock()
}

func GetTotalBalance() int {
 globalLock.Lock()
 total := 0
 for i := 0; i &lt; len(accounts); i++ {
  total += accounts[i].Balance()
 }
 globalLock.Unlock()
 return total
}

func RandomTranfer() {
 var sender, balance int
 for {
  sender = rand.Intn(len(accounts))
  balance = accounts[sender].Balance()
  if balance &gt; 0 {
   break
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
  RandomTranfer()
 }
}

func PrintTotalBalance() {
 fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
 for i := 0; i &lt; 20; i++ {
  accounts = append(accounts, &amp;Account{balance: 1000, mutex: &amp;sync.Mutex{}})
 }
 globalLock = &amp;sync.Mutex{}

 PrintTotalBalance()

 for i := 0; i &lt; 10; i++ {
  go GoTransfer()
 }

 for {
  PrintTotalBalance()
  time.Sleep(100 * time.Millisecond)
 }
}