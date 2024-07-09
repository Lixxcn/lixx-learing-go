package main

import "fmt"

func main() {
	var a, b int
	a, b = 1, 2
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
