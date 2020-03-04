package str

import "testing"

func TestStrlen(t *testing.T) {
	str := "Hello World"
	expected := 11
	result := strlen(str)
	if result != expected {
		t.Errorf("expected: %d, got: %d", expected, result)
	}

	str = "你好 世界"
	expected = 5
	result = strlen(str)
	if result != expected {
		t.Errorf("expected: %d, got: %d", expected, result)
	}
}
