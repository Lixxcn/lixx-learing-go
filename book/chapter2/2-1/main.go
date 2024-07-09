package main

import "fmt"

func main() {
	var a, b int
	a, b = 1, 2
	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
