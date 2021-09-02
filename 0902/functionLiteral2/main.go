package main

import "fmt"

func main() {

	i := 0
	fmt.Println(i)

	f := func() { //외부 변수를 내부로 가져오는 것 capture
		i += 10
	}

	i++
	fmt.Println(i)

	f()
	fmt.Println(i)
}
