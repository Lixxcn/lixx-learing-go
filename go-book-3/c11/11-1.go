package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	var num int
	var object interface{}
	fmt.Println("num", reflect.TypeOf(num), "kind", reflect.TypeOf(num).Kind())
	fmt.Println("object", reflect.TypeOf(object), "value", object)
	object = User{}
	fmt.Println("object", reflect.TypeOf(object), "value", object, "kind", reflect.TypeOf(object).Kind())
	object = User{
		Id:   2,
		Name: "Lixx",
		Age:  28,
	}
	fmt.Println("object", reflect.TypeOf(object), "value", object, "kind", reflect.TypeOf(object).Kind())

}
