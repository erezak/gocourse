package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
	"github.com/erezak/gocourse/c10sort/c10sort"
)

func main() {

	// Set goroutines pool to number of cores
	const numberOfTimes int = 10
	var numberOfItems int = int(math.Pow10(6))

	// By default, Go doesn't run the goroutines in parallel
	// You can either use code, or an environment variable
	// This should be changed once the Go scheduler is improved
	var numberOfCores int = runtime.NumCPU()
	runtime.GOMAXPROCS(numberOfCores)
	fmt.Printf("Using %d goroutines\n", numberOfCores)

	// Create random numbers slice
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := r.Perm(numberOfItems)

	fmt.Println(numbers[:20])
	sorted, timePassed := c10sort.Quicksort(numbers, true)
	fmt.Println(sorted[:20])
	fmt.Printf("sorted %d items in %d ms\n", numberOfItems, timePassed)
}
