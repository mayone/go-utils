package threadpool

import (
	"runtime"
	"sync"
)

type Runnable interface {
	Run(tInfo ThreadInfo)
}

type ThreadInfo struct {
	Id         int
	NumThreads int
}

type threadsPool struct {
	numThreads int
}

func NewThreadPool() *threadsPool {
	tp := threadsPool{
		numThreads: runtime.NumCPU(),
	}

	return &tp
}

func (tp threadsPool) worker(wg *sync.WaitGroup, id int, task Runnable) {
	defer wg.Done()
	threadInfo := ThreadInfo{
		Id:         id,
		NumThreads: tp.numThreads,
	}
	task.Run(threadInfo)
}

func (tp threadsPool) RunWorkers(task RunnableTask) {
	var wg sync.WaitGroup
	for id := 0; id < tp.numThreads; id++ {
		wg.Add(1)
		go tp.worker(&wg, id, task)
	}
	wg.Wait()
}
