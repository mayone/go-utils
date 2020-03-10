package str

import (
	"testing"
)

// func TestStrlen(t *testing.T) {
// 	assert := assert.New(t)

// 	str := "Hello World"
// 	expected := 11
// 	result := strlen(str)
// 	assert.Equal(expected, result)

// 	str = "你好 世界"
// 	expected = 5
// 	result = strlen(str)
// 	assert.Equal(expected, result)
// }

func Test_strlen(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "String with English characters",
			args: args{
				str: "Hello World",
			},
			want: 11,
		},
		{
			name: "String with Chinese characters",
			args: args{
				str: "你好 世界",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strlen(tt.args.str); got != tt.want {
				t.Errorf("strlen() = %v, want %v", got, tt.want)
			}
		})
	}
}
