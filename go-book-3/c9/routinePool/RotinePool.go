package routinepool

import "runtime"

type RoutinePool struct {
	coreNum  int
	TaskList chan Task
}

func (p *RoutinePool) Submit(t Task) {
	p.TaskList <- t
}

func NewPool(core int) *RoutinePool {
	if core == 0 {
		core = runtime.NumCPU()
	}

	p := &RoutinePool{
		TaskList: make(chan Task),
	}

	for i := 0; i < core; i++ {
		go func() {
			for task := range p.TaskList {
				task.Execute()
			}
		}()
	}

	return p
}
