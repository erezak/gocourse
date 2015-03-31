package main

import "fmt"

func main() {
	var nums [5]int = [5]int{1, 2, 3, 4, 5}
	
	for _,i := range nums {
		defer fmt.Println(i)
	}
	
	fmt.Println("regular line")
}