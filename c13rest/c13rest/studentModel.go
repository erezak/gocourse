package c13rest

import (
	"time"
)

type Student struct {
	StudentId   string
	Name        string
	DateOfBirth time.Time
}
