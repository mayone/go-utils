package sumofarray

import (
	"math/rand"
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

// func TestSumArray(t *testing.T) {
// 	a := []int{}
// 	expected := sumOfArraySequential(a)
// 	result := sumOfArrayParallel(a)
// 	if result != expected {
// 		t.Errorf("expected: %v, got: %v", expected, result)
// 	}

// 	a = randArray(1000)
// 	expected = sumOfArraySequential(a)
// 	result = sumOfArrayParallel(a)
// 	if result != expected {
// 		t.Errorf("expected: %v, got: %v", expected, result)
// 	}
// }

func Test_sumOfArray(t *testing.T) {
	a1000 := randArray(1000)
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Sum of empty array",
			args: args{
				a: []int{},
			},
			want: sumOfArraySequential([]int{}),
		},
		{
			name: "Sum of array with length 1000",
			args: args{
				a: a1000,
			},
			want: sumOfArraySequential(a1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfArrayParallel(tt.args.a); got != tt.want {
				t.Errorf("sumOfArrayParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}
