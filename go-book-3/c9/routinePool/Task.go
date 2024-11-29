package routinepool

type Task interface {
	Execute()
	GetResult() string
}
