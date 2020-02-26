package threadpool

import (
	"runtime"
	"sync"
)

type Runnable interface {
	Run(tInfo ThreadInfo)
}

// ThreadInfo is info of thread to be used when running task
type ThreadInfo struct {
	ID         int
	NumThreads int
}

type threadsPool struct {
	numThreads int
}

// NewThreadPool returns address of new threadPool
func NewThreadPool() *threadsPool {
	tp := threadsPool{
		numThreads: runtime.NumCPU(),
	}

	return &tp
}

func (tp threadsPool) worker(wg *sync.WaitGroup, id int, task Runnable) {
	defer wg.Done()
	threadInfo := ThreadInfo{
		ID:         id,
		NumThreads: tp.numThreads,
	}
	task.Run(threadInfo)
}

func (tp threadsPool) RunWorkers(task Runnable) {
	var wg sync.WaitGroup
	for id := 0; id < tp.numThreads; id++ {
		wg.Add(1)
		go tp.worker(&wg, id, task)
	}
	wg.Wait()
}
