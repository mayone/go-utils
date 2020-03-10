package addarrays

import "github.com/mayone/Go-Utilities/multithreading/threadpool"

type task struct {
	a []int
	b []int
	c []int
}

func (t *task) Run(tInfo threadpool.ThreadInfo) {
	offset := tInfo.ID
	step := tInfo.NumThreads
	for i := offset; i < len(t.c); i += step {
		t.c[i] = t.a[i] + t.b[i]
	}
}

func addArraysParallel(a, b []int) []int {
	if len(a) != len(b) {
		return nil
	}
	c := make([]int, len(a))
	tp := threadpool.NewThreadPool()
	task := &task{
		a: a,
		b: b,
		c: c,
	}
	tp.RunWorkers(task)

	return c
}

func addArraysSequential(a, b []int) []int {
	if len(a) != len(b) {
		return nil
	}
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] + b[i]
	}
	return c
}
