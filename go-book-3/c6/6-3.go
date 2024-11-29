package main

import (
	"fmt"
	"math"
)

func getMaxMin(sl *[]int) (max, min int) {
	max = math.MinInt64
	min = math.MaxInt64

	for _, e := range *sl {
		if e > max {
			max = e
		}
		if e < min {
			min = e
		}
	}
	fmt.Printf("max %p\n", &max)
	fmt.Printf("min %p\n", &min)
	return
}

func main() {
	sl := []int{3, 453, 423, 23, 543, -234}

	max, min := getMaxMin(&sl)

	fmt.Println(max, min)
	fmt.Printf("max %p\n", &max)
	fmt.Printf("min %p\n", &min)
}
