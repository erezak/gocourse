package c05interview

import "fmt"

type Manager struct {
	Employee
	EmployeesCount int
}

func (m Manager) BeInterviewed() string {
	return fmt.Sprintf("Hello, my name is %s and I'm a %s, with %d employees working for me",
		m.name, m.Position, m.EmployeesCount)
}
