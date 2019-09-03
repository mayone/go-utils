package swap

import "testing"

func TestSwap(t *testing.T) {
	a := 3
	b := 5
	Swap(&a, &b)
	expected := (a == 5 && b == 3)
	if !expected {
		t.Errorf("Swapping failed")
	}
}
