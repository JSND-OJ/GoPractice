package main

import "fmt"

type User struct {
	Name  string
	ID    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"오정헌", "OJ", 25, 5}
	vip := VIPUser{
		User{"너", "you", 30, 20},
		30,
		40,
	}
	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP유저: %s ID: %s 나이: %d VIP레벨: %d 유저레벨 %d\n",
		vip.Name,
		vip.ID,
		vip.Age,
		vip.Level,
		vip.User.Level,
	)
}
