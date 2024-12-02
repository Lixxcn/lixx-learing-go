package worker

import "fmt"

type Worker struct {
	Name     string
	Age      int
	Position string
}

func (worker Worker) Do() {
	fmt.Println(worker.Name, "正在工作")
	worker.Name = "xixi"
}
