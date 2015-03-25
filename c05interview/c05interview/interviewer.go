package c05interview

import "fmt"

type Interviewer Employee


// Note, Go doesn't support function overloading or default parameters
func (i Interviewer) Interview(employee Employee) string {
	return i.InterviewWithAdditionalInfo(employee, nil)
}

func (i Interviewer) InterviewWithAdditionalInfo(employee Employee, additionalInfo func() string) string {
	summary := fmt.Sprintf("%s has interviewed %s, who said: %s", i.name, employee.name, employee.BeInterviewed())
	if additionalInfo != nil {
		summary += fmt.Sprintf(" - %s", additionalInfo())
	}
	return summary
}