package main

import (
	"fmt"
	c "github.com/erezak/gocourse/c07sort/c07sort"
)

type myArray []int

func (m myArray) IsFirstLessThanSecondByIndex(firstIndex int, secondIndex int) bool {
	return m[firstIndex] < m[secondIndex]
}

func (m myArray) Length() int {
	return len(m)
}

func (m myArray) Swap(firstIndex int, secondIndex int) {
	m[firstIndex], m[secondIndex] = m[secondIndex], m[firstIndex]
}

func main() {
	var testArray = myArray{21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10,
		10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1}
	fmt.Printf("Original array: %v\n", testArray)
	rounds := c.BubbleSort(testArray)
	if rounds == 1 {
		fmt.Println("Original array was sorted")
	} else {
		fmt.Printf("\nSorted array (after %d rounds) is: %v\n", rounds, testArray)
	}
}
