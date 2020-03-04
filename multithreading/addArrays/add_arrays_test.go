package addarrays

import (
	"reflect"
	"testing"
)

func Test_addArraysParallel(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Add empty arrays",
			args: args{
				a: []int{},
				b: []int{},
			},
			want: addArraysSequential([]int{}, []int{}),
		},
		{
			name: "Add arrays with different length",
			args: args{
				a: []int{},
				b: []int{-1, -1, -1, -1, -1},
			},
			want: addArraysSequential([]int{}, []int{-1, -1, -1, -1, -1}),
		},
		{
			name: "Add arrays with same length",
			args: args{
				a: []int{1, 0, 2, 0, 3},
				b: []int{-1, -1, -1, -1, -1},
			},
			want: addArraysSequential([]int{1, 0, 2, 0, 3}, []int{-1, -1, -1, -1, -1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addArraysParallel(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addArraysParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}
