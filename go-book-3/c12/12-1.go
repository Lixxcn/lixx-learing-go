package main

import (
	"fmt"
	"reflect"
)

func sum[T float32 | float64 | int | string](a T, b T) T {
	fmt.Println("typt a b: ", reflect.TypeOf(a), reflect.TypeOf(b))
	return a + b
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1.1, 2.2))
	fmt.Println(sum("hello", "world"))
}
