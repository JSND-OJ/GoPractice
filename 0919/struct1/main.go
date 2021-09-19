package main

import (
	"fmt"
	"gopractice/0919//struct1/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "aaa"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}
