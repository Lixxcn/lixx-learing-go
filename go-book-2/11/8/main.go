package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println("goroutine")
		fmt.Println(v, v2)
	}()
	v := 2
	var v2 int
	time.Sleep(1 * time.Second)
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println("main goroutine")
	fmt.Println(v, v2)
	time.Sleep(1 * time.Second)
}
