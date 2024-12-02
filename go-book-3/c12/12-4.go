// 泛型类型提取为单独的类型定义
package main

import (
	"fmt"
	"reflect"
)

type SumPara interface{
	float32 | float64 | int | string
}

func sum[T SumPara](a T, b T) T {
	fmt.Println("typt a b: ", reflect.TypeOf(a), reflect.TypeOf(b))
	return a + b
}

func main() {
	fmt.Println(sum[int](1, 2))
	fmt.Println(sum[string]("hello", "world"))
}
