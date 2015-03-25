package main

import (
	"fmt"
	"errors"
)

type Calculation struct {
	operand1 int
	operator string
	operand2 int
}

func main() {
	var results []int
	var operatorUsage = map[string]int{"+": 0, "-": 0, "*": 0, "/": 0}
	var calculation = Calculation{}
	
	for {
		fmt.Print("Enter calcualtion: ")
		fmt.Scanf("%d %s %d\n", &calculation.operand1, &calculation.operator, &calculation.operand2)
		if calculation.operator == "@" {
			break
		}
		
		var result int
		var err error
		
		result, err = calculate(calculation)
		if err != nil {
			fmt.Println(err)
			continue
		}
		
		// Check if it's an calculation.operator - if it exists in the map
		if _, exists := operatorUsage[calculation.operator]; exists == true {
			operatorUsage[calculation.operator]++
		}
		
		results = append(results, result)
		
		fmt.Printf("%d %s %d = %d\n", calculation.operand1, calculation.operator, calculation.operand2, result)
	}
	fmt.Println()
	fmt.Println("Thank you for using the calculator\n\n")
	fmt.Printf("Operator usage: %v\n\n", operatorUsage)
	fmt.Printf("Results: %v\n\n", results)
	
}

func calculate(calculation Calculation) (result int, err error) {
	switch calculation.operator {
		case "+":
			result = calculation.operand1 + calculation.operand2
		case "*":
			result = calculation.operand1 * calculation.operand2
		case "-":
			result = calculation.operand1 - calculation.operand2
		case "/":
			if calculation.operand2 == 0 {
				err = errors.New("Please don't try to divide by 0")
			} else {
				result = calculation.operand1 / calculation.operand2
			}
		default:
			err = errors.New("%s is not a recognized calculation.operator\n")
	}
	return result, err
}