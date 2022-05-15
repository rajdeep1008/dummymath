// Package dummymath provides no use other than helping me play around with golang documentation features.
//
// Another paragraph of not any useful information.
package dummymath

// Sum adds all the numbers passed as arguments.
func Sum(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
