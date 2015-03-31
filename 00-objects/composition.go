package main

import "fmt"

type Name struct {
	first  string
	second string
}

type Person struct {
	id   string
	name Name
	age  int
}

type Employee struct {
	position string
	Person
}

func main() {
	var employee Employee
	var anotherPerson = Person{
		id: "021251486",
		name: Name{
			first:  "Abe",
			second: "Honesty"},
		age: 33}

	employee.age = 34
	employee.position = "CEO"

	if employee.age > anotherPerson.age {
		fmt.Println("he's older.")
	}
	
	fmt.
}
