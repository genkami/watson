package math

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	assertEq := func(actual, expected int64) {
		if actual != expected {
			t.Errorf("expected %d but got %d", expected, actual)
		}
	}
	var acc int64 = 1
	for i := int64(0); i < 15; i++ {
		assertEq(Pow(3, i), acc)
		acc = acc * 3
	}
}

func TestPowf(t *testing.T) {
	assertEq := func(actual, expected float64) {
		if math.Abs(expected-actual) > 1e-5 {
			t.Errorf("expected %f but got %f", expected, actual)
		}
	}
	var acc float64 = 1
	for i := int64(0); i < 15; i++ {
		assertEq(Powf(1.23, i), acc)
		acc = acc * 1.23
	}
}
