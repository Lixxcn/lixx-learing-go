package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	ch1 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i)
		}
		ch1 <- 0
	}()

	go func() {
		<-ch1
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i * 10)
		}
		ch <- 0
	}()

	<-ch
}
