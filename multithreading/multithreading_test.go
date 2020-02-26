package multithreading

import (
	"reflect"
	"testing"
)

func TestAddArrays(t *testing.T) {
	a := []int{}
	b := []int{}
	expected := addArraysSequential(a, b)
	result := addArraysParallel(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}

	a = []int{}
	b = []int{-1, -1, -1, -1, -1}
	expected = addArraysSequential(a, b)
	result = addArraysParallel(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}

	a = []int{1, 0, 2, 0, 3}
	b = []int{-1, -1, -1, -1, -1}
	expected = addArraysSequential(a, b)
	result = addArraysParallel(a, b)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
