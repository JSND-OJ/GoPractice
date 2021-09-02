package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 소지금액 입력 받는 기능
	fmt.Println("소지금액 입력.")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	cost, _ := strconv.Atoi(str)

	// 메뉴 보여주는 기능
	fmt.Println("------------------------------")
	fmt.Printf("\t\t1. 콜라 500원 2. 사이다 500원 3. 맥주 600원\n")
	fmt.Println("------------------------------")
	fmt.Println("마 함 골라봐라 : ")

	// 음료수 입력하는 기능
	res := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		drink, _ := reader.ReadString('\n')
		drink = strings.TrimSpace(drink)
		result, _ := strconv.Atoi(drink)
		res = result
		break
	}

	// 음료수 읽는 기능
	switch res {
	case 1:
		cost = cost - 500
		fmt.Println("콜라여.")
	case 2:
		cost = cost - 500
		fmt.Println("사이다여.")
	case 3:
		cost = cost - 600
		fmt.Println("맥주여.")
	default:
		cost = cost - 1000
		fmt.Println("아무거나 줄겡.")
	}

	if cost < 0 {
		fmt.Println("응 가고")
	} else {
		fmt.Printf("%v번 음료수가 나왔고, 거스름돈은 %v 입니다.", res, cost)
	}
}
