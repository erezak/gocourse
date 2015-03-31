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

func main() {
	var person Person
	var anotherPerson = Person{
		id: "021251486",
		name: Name{
			first:  "Abe",
			second: "Honesty"},
		age: 33}

	person.age = 34

	if person.age > anotherPerson.age {
		fmt.Println("he's older.")
	}
}
