// Homework-6: Standard library - Part 2
// Exercise 1 - Statistic extending tests
// Author: Eugene Klimov
// Date: 14 may 2019
package statistic

import "testing"

// Common struct for all math tests
type testPairs struct {
	values []float64
	result float64
}

// Average values
var testAver = []testPairs{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 2, 3, 4, 5, 6}, 3.5},
	{[]float64{-5.11, -6.12, -8.13}, -6.453333333333333},
}

// Sum values
var testSum = []testPairs{
	{[]float64{111.222, 333.444}, 444.666},
	{[]float64{1.1, 2.2, 3.3, 4.3, 5.3, 6.6}, 22.799999999999997},
	{[]float64{-10.1, -20.2, 30.3}, 3.552713678800501e-15},
}

// Main func for all Math tests
func TestMath(t *testing.T) {
	testSet(t, testAver, Average, "Average")
	testSet(t, testSum, Sum, "Sum")
}

// Common func for all Math calls
func testSet(t *testing.T, pairs []testPairs, f func([]float64) float64, name string) {
	for _, pair := range pairs {
		r := f(pair.values)
		if r != pair.result {
			t.Error(
				"In func", name,
				"for", pair.values,
				"expected", pair.result,
				"got", r,
			)
		}
	}
}
