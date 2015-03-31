package calculator

import (
	"testing"
)

// test steps convention:
// start with Test
// take a *testing.T
func TestAdd(t *testing.T) {
	var calc Calculator
	var expectedAnswer = 15.0
	if calc.Add(7.0, 8.0) != expectedAnswer {
		t.Error("Didn't work")
	}
}

func TestAddPrecise(t *testing.T) {
	var calc Calculator
	var expectedAnswer = 15.2
	if calc.Add(7.1, 8.1) != expectedAnswer {
		t.Error("Didn't work")
	}
}

func TestCompound(t* testing.T) {
	var calc Calculator
	var expectedAnswer = 15.2
	if calc.Substract(calc.Add(2.0, calc.Add(7.1, 8.1)),2.0) != expectedAnswer {
		t.Error("Didn't work")
	}
}

func TestDivideShouldNotReturnError(t *testing.T) {
	var calc Calculator
	_, err := calc.Divide(5, 2)
	if err != nil {
		t.Error("Divide by 0 didn't return an error")
	}

}

func TestDivideByZero(t *testing.T) {
	var calc Calculator
	_, err := calc.Divide(5, 0)
	if err == nil {
		t.Error("Divide by 0 didn't return an error")
	}

}
