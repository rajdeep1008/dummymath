package dummymath

import (
	"testing"
)

// TestSum tests the Sum function.
func TestSum(t *testing.T) {
	li := []int{1, 2, 3, 4, 5}

	want := 15
	got := Sum(li...)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
