package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i)
		}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond)
			fmt.Println(i * 10)
		}
	}()

	time.Sleep(time.Second)
}
