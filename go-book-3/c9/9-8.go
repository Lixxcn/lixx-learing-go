package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	ch := make(chan int, 1)

	go func() {

		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i)
		}
		wg.Done()
		ch <- 0
	}()

	go func() {
		<-ch
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i * 10)
		}
		wg.Done()
	}()

	wg.Wait()
}
