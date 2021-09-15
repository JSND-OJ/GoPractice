package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// txt 열기
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scanner로 읽기
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 파일 닫기, 에러 처리
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	//파일 읽는 도중 에러 처리
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
