package main

import "fmt"

type Stringer interface { //Stringer 인터페이스 선언. 매개변수 없이 string 타입을 반환하는 String메서드 포함
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string { //Student 타입은 String()메서드를 포함->Student 타입은 Stringer 인터페이스로 사용될 수 있음
	return fmt.Sprintf("Hi my name is %s.I'm %d years old", s.Name, s.Age) //문자열을 string 타입으로 반환
}

func main() {
	student := Student{"Jhon", 15}
	var stringer Stringer

	stringer = student //stringer 값으로 Student 타입의 변수 student 대입.

	fmt.Printf("%s\n", stringer.String()) //stringer 인터페이스가 가이고 있는 String메소드를 호출. stringer값으로 Student타입 student를 가지고 있기 때문에
}
