package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var operator string
	var results []int
	var operatorUsage = map[string]int{"+": 0, "-": 0, "*": 0, "/": 0}

	for {
		fmt.Print("Enter calcualtion: ")
		fmt.Scanf("%d %s %d\n", &num1, &operator, &num2)
		if operator == "@" {
			break
		}

		var result int
		
		switch operator {
			case "+":
				result = num1 + num2
			case "*":
				result = num1 * num2
			case "-":
				result = num1 - num2
			case "/":
				if num2 == 0 {
					fmt.Println("Please don't try to divide by 0")
					continue
				} else {
					result = num1 / num2
				}
			default:
				fmt.Printf("%s is not a recognized operator\n")
				continue
		}
		
		// Check if it's an operator - if it exists in the map
		if _, exists := operatorUsage[operator]; exists == true {
			operatorUsage[operator]++
		}
		
		results = append(results, result)
		
		fmt.Printf("%d %s %d = %d\n", num1, operator, num2, result)
	}
	fmt.Println()
	fmt.Println("Thank you for using the calculator\n\n")
	fmt.Printf("Operator usage: %v\n\n", operatorUsage)
	fmt.Printf("Results: %v\n\n", results)
	
}