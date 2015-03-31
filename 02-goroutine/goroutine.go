package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printNumber(number int, randomGenerator *rand.Rand) {
	timeToWait := time.Duration((randomGenerator.Int() % 3000) + 1000)
	time.Sleep(timeToWait * time.Millisecond)
	fmt.Printf("%3d\tWaited %d ms\n", number, timeToWait)
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 10; i++ {
		printNumber(i, r)
	}
	fmt.Scanln()
}
