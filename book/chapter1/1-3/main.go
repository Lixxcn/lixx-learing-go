package main

import (
	"flag"
	"fmt"
)

// go run main.go -intVal=27 -boolVal -strVal=Lixx
// go run main.go -intVal=27 -boolVal=true -strVal=Lixx
// go run main.go -intVal 27 -boolVal=false -strVal Lixx
func main() {
	// 定义一个整型标志位intVal，默认值为0，用于接收命令行参数-intVal
	// "int类型参数"是对intVal的描述，用于flag包的使用说明
	var intVal = flag.Int("intVal", 0, "int类型参数")

	var strVal = flag.String("strVal", "", "string类型参数")

	var boolVal = flag.Bool("boolVal", false, "bool类型参数")

	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数的值
	fmt.Println("-intVal:", *intVal)
	fmt.Println("-strVal:", *strVal)
	fmt.Println("-boolVal:", *boolVal)
}
