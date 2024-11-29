package main

import (
	"fmt"
	"sync"
	"time"
)

type SingleObject struct {
}

var once sync.Once

var instance *SingleObject

func getInstance() *SingleObject {
	once.Do(func() {
		instance = &SingleObject{}
		fmt.Println("实例化SingleObject")
	})
	return instance
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			_ = getInstance()
		}()
	}

	time.Sleep(time.Second)
}
