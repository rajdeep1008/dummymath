package dummymath

import (
	"fmt"
	"testing"
)

type testpair struct {
	values []int
	sum    int
}

var tests = []testpair{
	{values: []int{1, 2, 3, 4, 5}, sum: 15},
	{values: []int{}, sum: 0},
	{values: []int{0}, sum: 0},
	{values: []int{-1, 1, 3, -3}, sum: 0},
	{values: []int{-1, -1, -3, -3}, sum: -8},
}

// TestSum tests the Sum function.
func TestSum(t *testing.T) {
	for _, pair := range tests {
		got := Sum(pair.values...)
		want := pair.sum

		if got != want {
			t.Error(
				"For", pair.values,
				"expected", pair.sum,
				"got", got,
			)
		}
	}
}

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3))
	// Output:
	// 6
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3)
	}
}
