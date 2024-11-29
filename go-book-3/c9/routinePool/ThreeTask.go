package routinepool

import "sync"

type DockerTack struct {
	Wg     *sync.WaitGroup
	Result string
}

func (task *DockerTack) Execute() {
	task.Result = "docker task"
	task.Wg.Done()
}

func (task *DockerTack) GetResult() string {
	return task.Result
}

type WorkerTack struct {
	Wg     *sync.WaitGroup
	Result string
}

func (task *WorkerTack) Execute() {
	task.Result = "worker task"
	task.Wg.Done()
}

func (task *WorkerTack) GetResult() string {
	return task.Result
}

type TeacherTack struct {
	Wg     *sync.WaitGroup
	Result string
}

func (task *TeacherTack) Execute() {
	task.Result = "teacher task"
	task.Wg.Done()
}

func (task *TeacherTack) GetResult() string {
	return task.Result
}
