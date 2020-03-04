package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrlen(t *testing.T) {
	assert := assert.New(t)

	str := "Hello World"
	expected := 11
	result := strlen(str)
	assert.Equal(expected, result)
	// if result != expected {
	// 	t.Errorf("expected: %d, got: %d", expected, result)
	// }

	str = "你好 世界"
	expected = 5
	result = strlen(str)
	assert.Equal(expected, result)
	// if result != expected {
	// 	t.Errorf("expected: %d, got: %d", expected, result)
	// }
}
