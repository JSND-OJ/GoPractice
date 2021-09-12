package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//타겟 난수 생성
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("1에서 100사이의 난수 생성 완료.")
	fmt.Println("뭐게?")

	reader := bufio.NewReader(os.Stdin)
	success := false

	for t := 0; t < 10; t++ {
		fmt.Printf("%d번 남음\n", 10-t)

		//입력
		fmt.Println("입력 해봐")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//입력 받은 문자열 int형으로 변환
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		//비교
		if guess < target {
			fmt.Println("너무 작다")
		} else if guess > target {
			fmt.Println("너무 크다")
		} else {
			success = true
			fmt.Println("정답!")
			fmt.Printf("%d = %d", target, guess)
			break
		}
		if !success {
			fmt.Println("어렵구만")
		}
	}
	if !success {
		fmt.Printf("몰름? %d임", target)
	}
}
