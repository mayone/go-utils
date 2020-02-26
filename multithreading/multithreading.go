package multithreading

import (
	"runtime"
	"sync"
)

// func worker(wg *sync.WaitGroup, work func()) {
// 	defer wg.Done()
// 	work()
// }

type workFunc func(id int)

type threadsPool struct {
	nCPU int
	work workFunc
}

func (tp threadsPool) worker(wg *sync.WaitGroup, id int, task Runnable) {
	defer wg.Done()
	task.Run(id, tp.nCPU)
}

func (tp threadsPool) RunWorkers(task Runnable) {
	var wg sync.WaitGroup
	for i := 0; i < tp.nCPU; i++ {
		wg.Add(1)
		go tp.worker(&wg, i, task)
	}
	wg.Wait()
}

func newThreadPool() *threadsPool {
	tp := threadsPool{
		nCPU: runtime.NumCPU(),
	}

	return &tp
}

type Runnable interface {
	Run(id int, nCPU int)
}

type task struct {
	a []int
	b []int
	c []int
}

func (t *task) Run(id int, nCPU int) {
	offset := id
	step := nCPU
	for i := offset; i < len(t.c); i += step {
		t.c[i] = t.a[i] + t.b[i]
	}
}

func addArraysParallel(a []int, b []int) []int {
	// nCPU := runtime.NumCPU()
	// var wg sync.WaitGroup

	// worker := func(wg *sync.WaitGroup, a []int, b []int, c []int, offset int, step int) {
	// 	defer wg.Done()
	// 	for i := offset; i < len(c); i += step {
	// 		c[i] = a[i] + b[i]
	// 	}
	// }

	// work := func(a []int, b []int, c []int, offset int, step int) {
	// 	for i := offset; i < len(c); i += step {
	// 		c[i] = a[i] + b[i]
	// 	}
	// }

	if len(a) != len(b) {
		return nil
	}
	c := make([]int, len(a))
	tp := newThreadPool()
	task := &task{a: a, b: b, c: c}
	tp.RunWorkers(task)
	// for i := 0; i < nCPU; i++ {
	// 	wg.Add(1)
	// 	go worker(&wg, a, b, c, i, nCPU)
	// }
	// wg.Wait()

	return c
}

func addArraysSequential(a []int, b []int) []int {
	if len(a) != len(b) {
		return nil
	}
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}
