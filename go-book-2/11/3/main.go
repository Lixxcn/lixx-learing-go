package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println("左手画圆")
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("右手画方")
		time.Sleep(1 * time.Second)
	}
}
