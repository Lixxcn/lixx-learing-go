package main

import "fmt"

func main() {
	i1 := cumulate(1, 2)
	formatOutput(i1)
	fmt.Println(i1)
}

func cumulate(i, total int) int {
	total = i + total
	return total
}

func formatOutput(i int) {
	fmt.Println(i)
}
