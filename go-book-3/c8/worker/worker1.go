package worker

import "fmt"

type Worker1 struct {
	Name     string
	Age      int
	Position string
}

func (worker *Worker1) Do1() {
	fmt.Println(worker.Name, "正在工作")
	worker.Name = "xixi"
}
