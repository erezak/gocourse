package main

import (
	"fmt"
	"github.com/erezak/gocourse/c09pomodoro/c09pomodoro"
)

func main() {
	timerOutputs := make(chan string)
	stopSingal := make(chan bool)
	timerDone := make(chan bool)
	var timer = c09pomodoro.NewTimer(5, 25, 5)
	fmt.Println("Starting timer - press enter to stop.")
	go timer.Start(timerOutputs, timerDone, stopSingal)

	go func() {
		fmt.Scanln()
		stopSingal <- true
	}()

	done := false
	for !done {
		select {
		case output := <-timerOutputs:
			fmt.Println(output)
		case <-timerDone:
			done = true
		}
	}

}
