package sumOfArray

import (
	"math/rand"
	"os"
	"runtime/trace"
	"testing"
	"time"
)

func randArray(len int) []int {
	a := make([]int, len)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		a[i] = rand.Intn(100)
	}
	return a
}

func TestSumArray(t *testing.T) {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	a := []int{}
	expected := sumOfArraySequential(a)
	result := sumOfArrayParallel(a)
	if result != expected {
		t.Errorf("expected: %v, got: %v", expected, result)
	}

	a = randArray(1000)
	expected = sumOfArraySequential(a)
	result = sumOfArrayParallel(a)
	if result != expected {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
