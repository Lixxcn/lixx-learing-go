package main

import (
	"fmt"
	"time"
)

func handle() {
	fmt.Println("begine")
	fmt.Println("end")
}

func main() {
	// go fmt.Println("启动新的协程")

	go handle()

	time.Sleep(time.Microsecond * 100)
}
