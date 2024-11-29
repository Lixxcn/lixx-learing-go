package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	lock := sync.RWMutex{}

	go func() {
		lock.RLock()
		defer lock.RUnlock()
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i)
		}
	}()

	go func() {
		lock.Lock()
		defer lock.Unlock()
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i * 10)
		}
	}()

	time.Sleep(time.Second)
}
