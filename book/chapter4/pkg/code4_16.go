package pkg

import (
	"fmt"
	"os"
	"time"
)

func DoRunning(messageChan chan int) {
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

func CancelRunning(messageChan chan int) {
	os.Stdin.Read(make([]byte, 1))
	messageChan <- 1
}
