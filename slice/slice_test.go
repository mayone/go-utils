package slice

import (
	"reflect"
	"testing"
)

// func TestRemoveByIndex(t *testing.T) {
// 	slice := []int{1, 2, 3, 4, 5}
// 	removeByIndex(&slice, 2)
// 	expected := []int{1, 2, 4, 5}
// 	result := slice
// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("expected: %v, got: %v", expected, result)
// 	}
// }

func Test_removeByIndex(t *testing.T) {
	type args struct {
		slice []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Remove head",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				index: 0,
			},
			want: []int{2, 3, 4, 5},
		},
		{
			name: "Remove middle",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				index: 2,
			},
			want: []int{1, 2, 4, 5},
		},
		{
			name: "Remove tail",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				index: 4,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Remove index out of bound",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				index: 5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeByIndex(tt.args.slice, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
