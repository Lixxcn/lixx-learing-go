package main

import (
	"fmt"
	"os"
	"time"
)

func doRunning(messageChan chan int) {
	c := time.Tick(1 * time.Second)

	for {
		select {
		case <-messageChan:
			fmt.Println("receive message, stop running.")
			return
		default:
			runningTime := <-c
			fmt.Printf("running time: %v\n", runningTime)
		}
	}
}

func cancelRunning(messageChan chan int) {
	os.Stdin.Read(make([]byte, 1))
	messageChan <- 1
}

func main() {
	messageChan := make(chan int)
	go doRunning(messageChan)
	go cancelRunning(messageChan)
	time.Sleep(100 * time.Second)
}
