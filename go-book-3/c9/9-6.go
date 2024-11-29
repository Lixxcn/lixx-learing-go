package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() {

		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i * 10)
		}
		wg.Done()
	}()

	wg.Wait()
}
