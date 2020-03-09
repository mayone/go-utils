package sumofarray

import (
	"github.com/mayone/Go-Utilities/multithreading/threadpool"
)

type task struct {
	a  []int
	ch chan int64
}

func (t *task) Run(tInfo threadpool.ThreadInfo) {
	offset := tInfo.ID
	step := tInfo.NumThreads
	var sum int64 = 0

	for i := offset; i < len(t.a); i += step {
		sum += int64(t.a[i])
	}
	select {
	case current := <-t.ch:
		t.ch <- current + sum
	}
}

func sumOfArrayParallel(a []int) int64 {
	if len(a) == 0 {
		return 0
	}
	var sum int64 = 0
	ch := make(chan int64, 1)
	ch <- sum

	tp := threadpool.NewThreadPool()
	task := &task{
		a:  a,
		ch: ch,
	}
	tp.RunWorkers(task)
	sum = <-ch

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
