package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var a int
	fmt.Printf("온도넣으세요-->>>")
	fmt.Scanln(&a)
	var b = [5]string{"순두부", "중국집", "국밥집", "집밥집", "신세계"}
	rand.Seed(time.Now().UnixNano())
	t := rand.Intn(5)

	if a < 25 {
		t = rand.Intn(5)
		fmt.Println("25도보다 낮네 250m내에 있는 신세계,집밥집,국밥집,중국집,순두부 중에 골라줄게--------->>", b[t])

	} else if a < 30 {
		t = rand.Intn(4)
		fmt.Println("25도~30도 사이네 200m내에 있는 집밥집,국밥집,중국집,순두부 중에 골라줄게--------->>", b[t])
	} else if a < 35 {

		t = rand.Intn(3)
		fmt.Println("30에서 35도사이네 150m내에 있는 국밥집,중국집,순두부 중에 골라줄게--------->>", b[t])
	} else if a < 40 {
		t = rand.Intn(2)
		fmt.Println("35에서 40도사이네 100m내에 있는 국밥집,중국집,순두부 중에 골라줄게--------->>", b[t])
	} else {
		fmt.Println("40도보다 높네 50m내에 있는 순두부 중에 골라줄게--------->>", "순두부")
	}
	//fmt.Println(a)
}
