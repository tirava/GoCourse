// Homework-6: Standard library - Part 2
// Exercise 4 - Square Equation Tests
// Author: Eugene Klimov
// Date: 15 may 2019
package square

import (
	"math"
	"testing"
)

type testValue struct {
	a, b, c float64
	d       float64 // Discriminant
	x1, x2  float64
}

// testValues for testing square equation
var testValues = []testValue{ // x1 is for minus calc in formula first, x2 is for plus in second time
	{1, 0, 0, 0, 0, 0},
	{1, 0, 1, -4, 0, 0},
	{1, 1, 0, 1, -1, 0},
	{-5.5, 5.4, 5.3, 145.76, 1.588464601657205, -0.6066464198390232}, // two roots
	{-13, 26, -13, 0, 1, 1},                           // one root
	{123.45, 67.89, 10.11, -383.26590000000033, 0, 0}, // no roots
}

// TestDiscriminant tests discriminant only
func TestDiscriminant(t *testing.T) {
	for _, td := range testValues {
		_, _, d := EquationSquare(td.a, td.b, td.c)
		if d != td.d {
			t.Error("\nFor a =", td.a, ", b =", td.b, ", c =", td.c,
				"expected discriminant =", td.d, "\nbut got", d)
		}
	}
}

// Test1stArg checks result if 'a=0' - equation impossible
func Test1stArg(t *testing.T) {
	x1, x2, d := EquationSquare(0, 0, 0)
	if !math.IsNaN(x1) || !math.IsNaN(x2) || !math.IsNaN(d) {
		t.Error("\nFor a =", 0,
			"expected 'x1 = NaN, x2 = NaN, d = NaN' (i.e. impossible)\n but got",
			"x1 =", x1, ", x2 =", x2, ", d =", d)
	}
}

// TestRoots checks real roots of square equation
func TestRoots(t *testing.T) {
	for _, tr := range testValues {
		x1, x2, _ := EquationSquare(tr.a, tr.b, tr.c)
		if x1 != tr.x1 || x2 != tr.x2 {
			t.Error("\nFor a =", tr.a, ", b =", tr.b, ", c =", tr.c,
				"expected x1 =", tr.x1, "x2 =", tr.x2, "\nbut got",
				"x1 =", x1, ", x2 =", x2)
		}
	}
}
