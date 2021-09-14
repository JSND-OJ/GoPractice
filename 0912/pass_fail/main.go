package main

import (
	"GoPractice/myPKG/input"
	"fmt"
	"log"
)

var status string

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := input.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
