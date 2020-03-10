package greet

import "testing"

func TestGreet(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "From Wayne",
			args: args{
				from: "Wayne",
				to:   "World",
			},
			want: "Hello World, I am Wayne",
		},
		{
			name: "From World",
			args: args{
				from: "World",
				to:   "Wayne",
			},
			want: "Hello Wayne, I am World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
