package threadpool

import (
	"runtime"
	"sync"
)

type ThreadInfo struct {
	Id         int
	NumThreads int
}

type threadsPool struct {
	numThreads int
}

func (tp threadsPool) worker(wg *sync.WaitGroup, id int, task Runnable) {
	defer wg.Done()
	threadInfo := ThreadInfo{Id: id, NumThreads: tp.numThreads}
	task.Run(threadInfo)
}

func (tp threadsPool) RunWorkers(task Runnable) {
	var wg sync.WaitGroup
	for i := 0; i < tp.numThreads; i++ {
		wg.Add(1)
		go tp.worker(&wg, i, task)
	}
	wg.Wait()
}

func NewThreadPool() *threadsPool {
	tp := threadsPool{
		numThreads: runtime.NumCPU(),
	}

	return &tp
}

type Runnable interface {
	Run(tInfo ThreadInfo)
}
