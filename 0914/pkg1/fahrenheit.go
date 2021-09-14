package main

import (
	"GoPractice/0914/input"
	"fmt"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	input.GetFloat()
	fahrenheit, err := input.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
