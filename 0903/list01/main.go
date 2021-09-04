package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New() //새로운 리스트 생성
	fmt.Println(v.Len())

	e4 := v.PushBack(4) //리스트 뒤에 요소 추가 [4]
	fmt.Println(v.Len())

	e1 := v.PushFront(1) //리스트 앞에 요소 추가 [1][4]
	fmt.Println(v.Len())

	v.InsertBefore(3, e4) //e4 요소 앞에 요소 삽입 [1][3][4]
	fmt.Println(v.Len())

	v.InsertAfter(2, e1) //e4 요소 뒤에 요소 삽입
	fmt.Println(v.Len())

	for e := v.Front(); e != nil; e = e.Next() { //초기문 e := v.Front() 실행 -->  조건문 e != nil 검사 --> 함수 본문 실행 --> 후처리 e = e.Next() 실행 --> 조건 --> 본문 --> 후처리 반복
		fmt.Print(e.Value, "")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, "")
	}
}
