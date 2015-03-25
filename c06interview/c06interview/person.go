package c06interview

type Person struct {
	id   string
	name string
	age  int
}

func NewPerson(id string, name string, age int) Person {
	return Person{id: id,
		name: name,
		age:  age}
}

func (p Person) Name() string {
	return p.name
}