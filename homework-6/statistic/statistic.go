// Homework-6: Standard library - Part 2
// Exercise 1 - Statistic extending
// Author: Eugene Klimov
// Date: 14 may 2019

package statistic

// Average is unchanged fom original author
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Sum calculates sum of the slice elements
func Sum(xs []float64) (sum float64) {
	for _, x := range xs {
		sum += x
	}
	return sum
}
