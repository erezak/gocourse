package main

import (
	"fmt"
	"github.com/erezak/gocourse/c04calculator/c04calculator"
)

func main() {

	// Note the use of the "constructor"
	// It is mandated by setting the calculator struct as private
	var calculator = c04calculator.NewCalculator()

	var operand1 int
	var operator string
	var operand2 int

	for {
		fmt.Print("Enter calcualtion: ")
		fmt.Scanf("%d %s %d\n", &operand1, &operator, &operand2)
		if operator == "@" {
			break
		}

		var err error

		calculator.SetCalculation(operand1, operator, operand2)

		err = calculator.Calculate()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("%v\n", calculator.CurrentCalculation())
	}
	fmt.Println()
	fmt.Println("Thank you for using the calculator\n\n")
	fmt.Printf("Operators usage: %v\n\n", calculator.OperatorsCount())
	fmt.Printf("Results: %+v\n\n", calculator.History())

}
