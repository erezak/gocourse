// Package used for demonstrating interfaces usage in Go
// as part of the OOP in Go section of the Go Course
package c06interview

import "fmt"

type Interviewer Employee

type BeIntervieweder interface {
	BeInterviewed() string
	Name() string
}

func (i Interviewer) Interview(employee BeIntervieweder) string {
	summary := fmt.Sprintf("%s has interviewed %s, who said: %s", i.name, employee.Name(), employee.BeInterviewed())
	return summary
}