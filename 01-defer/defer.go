package main

import "fmt"

func main() {
	defer fmt.Println("defered line")
	fmt.Println("regular line")
}