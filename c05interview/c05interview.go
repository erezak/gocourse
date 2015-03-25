package main

import (
	"fmt"
	c "github.com/erezak/gocourse/c05interview/c05interview"
)

func main() {
	var employee = c.Employee{Position: "Secretary",
		Person: c.NewPerson("029648912", "Billy Clint", 39)}
	var interviewer = c.Interviewer{Position: "Interviewerr",
		Person: c.NewPerson("547125687", "Sarah Moses", 51)}

	var manager = c.Manager{Employee: c.Employee{Position: "Manager",
		Person: c.NewPerson("547125687", "Amy Sil", 26)},
		EmployeesCount: 17}

	fmt.Println(interviewer.Interview(employee))
	fmt.Println()
	fmt.Println(manager)
	// The following line will not compile
	fmt.Println(interviewer.Interview(manager.Employee))
}
