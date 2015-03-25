package c06interview

type Employee struct {
	Position string
	Person
}

func (e Employee) BeInterviewed() string {
	return "Hello, my name is " + e.name + " and I'm a " + e.Position
}
