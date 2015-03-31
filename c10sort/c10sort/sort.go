package c10sort

import (
	"time"
	"runtime"
)

var runInParllel bool

func Quicksort(nums []int, parallel bool) ([]int, int) {
	started := time.Now()
	ch := make(chan int)
	runInParllel = parallel
	
	go quicksort(nums, ch)
	
	sorted := make([]int, len(nums))
	i := 0
	for next := range ch {
		sorted[i] = next
		i++
	}
	return sorted, int(time.Since(started).Nanoseconds() / 1000000)
}

func quicksort(nums []int, ch chan int) {

	// Choose first number as pivot
	pivot := nums[0]

	// Prepare secondary slices
	smallerThanPivot := make([]int, 0)
	largerThanPivot := make([]int, 0)

	// Slice except pivot
	nums = nums[1:]

	// Go over slice and sort
	for _, i := range nums {
		switch {
		case i <= pivot:
			smallerThanPivot = append(smallerThanPivot, i)
		case i > pivot:
			largerThanPivot = append(largerThanPivot, i)
		}
	}
	
	var ch1 chan int
	var ch2 chan int
	
	// Now do the same for the two slices
	if len(smallerThanPivot) > 1 {
		ch1 = make(chan int, len(smallerThanPivot))
		if runInParllel  && runtime.NumGoroutine() < runtime.GOMAXPROCS(0){
			go quicksort(smallerThanPivot, ch1)
		} else {
			quicksort(smallerThanPivot, ch1)
		}
	}
	if len(largerThanPivot) > 1 {
		ch2 = make(chan int, len(largerThanPivot))
		if runInParllel  && runtime.NumGoroutine() < runtime.GOMAXPROCS(0){
			go quicksort(largerThanPivot, ch2)
		} else {
			quicksort(largerThanPivot, ch2)
		}
	}

	// Wait until the sorting finishes for the smaller slice
	if len(smallerThanPivot) > 1 {
		for i := range ch1 {
			ch <- i
		}
	} else if len(smallerThanPivot) == 1 {
		ch <- smallerThanPivot[0]
	}
	ch <- pivot

	if len(largerThanPivot) > 1 {
		for i := range ch2 {
			ch <- i
		}
	} else if len(largerThanPivot) == 1 {
		ch <- largerThanPivot[0]
	}

	close(ch)
}
