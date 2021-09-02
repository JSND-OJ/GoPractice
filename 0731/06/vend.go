//1. 소유 금액 입력.
//2. 자판기 메뉴 출력
//3. 번호로 메뉴 입력, 소지금액 차감
//4. 메뉴가 나왔다는 문구, 이후 무한반복
//5. 소지금액 소진 시 탈출

package main

import "fmt"

func main() {

	// 소지금액 입력을 위해 소지금액 받는 변수 설절
	var budget int
	for {
		fmt.Println("소지금액 입력")
		_, error := fmt.Scanf("%d\n", &budget)

		if error != nil {
			fmt.Println("잘 못 입력했습니다.")
		} else {
			if (cost < 3000) || (cost > 4000) {
				continue
			} else {
				break
			}
		}
	}
}
