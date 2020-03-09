package sumofarray

import (
	"sync/atomic"

	"github.com/mayone/Go-Utilities/multithreading/threadpool"
)

type task struct {
	a   []int
	sum *int64
}

func (t *task) Run(tInfo threadpool.ThreadInfo) {
	offset := tInfo.ID
	step := tInfo.NumThreads
	var sum int64 = 0

	for i := offset; i < len(t.a); i += step {
		sum += int64(t.a[i])
	}
	atomic.AddInt64(t.sum, sum)
	// *t.sum += sum
}

func sumOfArrayParallel(a []int) int64 {
	if len(a) == 0 {
		return 0
	}
	var sum int64 = 0

	tp := threadpool.NewThreadPool()
	task := &task{
		a:   a,
		sum: &sum,
	}
	tp.RunWorkers(task)

	return sum
}

func sumOfArraySequential(a []int) int64 {
	if len(a) == 0 {
		return 0
	}
	var sum int64 = 0
	for _, n := range a {
		sum += int64(n)
	}
	return sum
}
