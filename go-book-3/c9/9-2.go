package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU核心数：", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	fmt.Println("GOMAXPROCS默认值：", runtime.GOMAXPROCS(0))
}
