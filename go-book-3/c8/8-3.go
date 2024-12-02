package main

import (
	"fmt"

	"github.com/Lixxcn/lixx-learning-go/go-book-3/c8/worker"
)

func main() {
	worker := &worker.Worker1{
		Name: "Lixx",
		Age:  27,
	}
	worker.Do1()
	fmt.Println(worker.Name)
}
