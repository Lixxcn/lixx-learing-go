package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Golang")
	fmt.Println("参数个数 = ", len(os.Args))

	for i, arg := range os.Args {
		fmt.Printf("参数%d = %s\n", i, arg)
	}
}
