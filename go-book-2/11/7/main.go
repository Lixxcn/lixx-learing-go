package main

import "fmt"

func main() {
	chel := make(chan int, 6)

	fmt.Println(len(chel))

	chel <- 0
	chel <- 1
	chel <- 2

	fmt.Println(len(chel))
}
