package c04calculator

import (
	"errors"
)

type Calculation struct {
	operand1 int
	operator string
	operand2 int
	result int
}

type calculator struct {
	currentCalculation Calculation
	history [] Calculation
	operatorsCount map[string]int
}

func NewCalculator() calculator{
	return calculator{operatorsCount: map[string]int{"+": 0,
													 "/": 0,
													 "*": 0,
													 "-": 0}}
}

// Public properties
func (c calculator) OperatorsCount() map[string]int {
	return c.operatorsCount
}

func (c calculator) CurrentCalculation() Calculation {
	return c.currentCalculation
}

func (c calculator) History() [] Calculation {
	return c.history
}

// Public methods
func (c *calculator) SetCalculation(operand1 int, operator string, operand2 int) {
	c.currentCalculation = Calculation {operand1: operand1,
										operator: operator,
										operand2: operand2}
}

func (c *calculator) Calculate () error {
	var err error
	switch c.currentCalculation.operator {
	case "+":
		c.currentCalculation.result = c.currentCalculation.operand1 + c.currentCalculation.operand2
	case "*":
		c.currentCalculation.result = c.currentCalculation.operand1 * c.currentCalculation.operand2
	case "-":
		c.currentCalculation.result = c.currentCalculation.operand1 - c.currentCalculation.operand2
	case "/":
		if c.currentCalculation.operand2 == 0 {
			err = errors.New("Please don't try to divide by 0")
		} else {
			c.currentCalculation.result = c.currentCalculation.operand1 / c.currentCalculation.operand2
		}
	default:
		err = errors.New("%s is not a recognized currentCalculation.operator\n")
	}
	
	// Add to operators map and to history
	if err == nil {
		c.operatorsCount[c.currentCalculation.operator]++
		c.history = append(c.history, c.currentCalculation)
	}
	
	return err
}
