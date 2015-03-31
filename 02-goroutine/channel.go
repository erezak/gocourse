package main

import (
	"fmt"
	"math/rand"
	"time"
)


func countChars(url string, randomGenerator *rand.Rand, count chan <- int) {
	timeToWait := time.Duration((randomGenerator.Int() % 3000) + 1000)
	charactersCount := randomGenerator.Int() % 20000
	time.Sleep(timeToWait * time.Millisecond)
	fmt.Printf("%s took %d ms\n", url, timeToWait)
	count <- charactersCount
}

func process(urls []string, randomGenerator *rand.Rand) int {
	total := 0
	charactersCount := make(chan int, len(urls))
	for _,url:= range urls {
		go countChars(url, randomGenerator, charactersCount)
	}
	for _= range urls {
		total += <- charactersCount
	}

	return total
}
func main() {
	startTime := time.Now()
	var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com"}
	total := process(urls, r)
	fmt.Printf("\n\nUrls have %d bytes in total.\nProcess took %d ms\n", total, time.Since(startTime)/1000000)
}
