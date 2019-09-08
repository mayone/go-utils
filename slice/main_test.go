package main

import (
	"reflect"
	"testing"
)

func TestRemoveByIndex(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	RemoveByIndex(&slice, 2)
	expected := []int{1, 2, 4, 5}
	result := slice
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
