package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)

	go func() {
		fmt.Println("子协程运行")
	}()

	runtime.Gosched()
}
