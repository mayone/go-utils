package sumOfArray

import (
	"github.com/mayone/OGGo/Go-Utilities/multithreading/threadpool"
)

type task struct {
	a   []int
	sum *int
}

func (t *task) Run(tInfo threadpool.ThreadInfo) {
	offset := tInfo.ID
	step := tInfo.NumThreads
	sum := 0
	for i := offset; i < len(t.a); i += step {
		sum += t.a[i]
	}
	*t.sum += sum
}

func sumOfArrayParallel(a []int) int {
	if len(a) == 0 {
		return 0
	}
	sum := 0

	tp := threadpool.NewThreadPool()
	task := &task{
		a:   a,
		sum: &sum,
	}
	tp.RunWorkers(task)

	return sum
}

func sumOfArraySequential(a []int) int {
	if len(a) == 0 {
		return 0
	}
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}
