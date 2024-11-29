package main

import "fmt"

func changeElement(arr [5]int) {
	arr[0] = 0
}

func changeElementByPointer(arr *[5]int) {
	arr[0] = 0
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	changeElement(arr)
	fmt.Println(arr[0])

	changeElementByPointer(&arr)
	fmt.Println(arr[0])
}
