// Homework-6: Standard library - Part 2
// Exercise 4 - Square Equation
// Author: Eugene Klimov
// Date: 15 may 2019
package square

import "math"

// EquationSquare returns roots and discriminant
func EquationSquare(a, b, c float64) (x1, x2, d float64) {
	if a == 0 { // check if 1st arg must not be = 0
		return 0, 0, 0
	}
	// calc discriminant
	d = b*b - 4*a*c
	if d == 0 { // one root
		x1 = -b / (2 * a) // no need calc common 'a2 := 2 * a' for best performance
		x2 = x1
	} else if d > 0 { // two roots
		a2 := 2 * a
		ds := math.Sqrt(d) // common a2 & ds for best performance
		x1 = (-b - ds) / a2
		x2 = (-b + ds) / a2
	} // if d < 0 then x1=x2=0
	return x1, x2, d
}
