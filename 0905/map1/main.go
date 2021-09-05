package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["오"] = "정헌"
	m["박"] = "아달라"

	m["오"] = "원빈"

	fmt.Printf("오씨의 이름은 %s입니다\n", m["오"])
	fmt.Printf("박씨의 이름은 %s이니다", m["박"])
}
