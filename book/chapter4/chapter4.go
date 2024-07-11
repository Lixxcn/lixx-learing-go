package main

import (
	"fmt"
	"lixx-learning-go/book/chapter4/pkg"
	"time"
)

func main() {
	fmt.Println("第四章")
	pkg.F4_15()
	// persion := pkg.Persion{}
	// persion.

	messageChan := make(chan int)
	go pkg.DoRunning(messageChan)
	go pkg.CancelRunning(messageChan)
	time.Sleep(100 * time.Second)
}
