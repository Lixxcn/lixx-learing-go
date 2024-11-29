package main

import (
	"fmt"
	"sync"

	routinepool "github.com/Lixxcn/lixx-learning-go/go-book-3/c9/routinePool"
)

func main() {
	pool := routinepool.NewPool(2)
	var wg sync.WaitGroup

	doctorTask := &routinepool.DockerTack{
		Wg: &wg,
	}
	wg.Add(1)
	pool.Submit(doctorTask)

	workerTask := &routinepool.WorkerTack{
		Wg: &wg,
	}
	wg.Add(1)
	pool.Submit(workerTask)

	teacherTask := &routinepool.TeacherTack{
		Wg: &wg,
	}
	wg.Add(1)
	pool.Submit(teacherTask)

	wg.Wait()

	fmt.Println(doctorTask.GetResult())
	fmt.Println(workerTask.GetResult())
	fmt.Println(teacherTask.GetResult())

}
