package main

import "testing"

func TestPrimePositiveCase(t *testing.T) {
	// These are valid two side prime numbers.
	numbers := []int{2, 3, 5, 7, 23, 37, 53, 73, 313, 317, 373, 797, 3137, 3797, 739397}

	for _, n := range numbers {
		expected := true

		actual := checkTwoSidePrime(n)

		if actual != expected {
			t.Errorf("Expected of %v to be %t but got %t", n, expected, actual)
		}
	}
}

func TestPrimeNegativeCase(t *testing.T) {
	// These are NOT valid two side prime numbers.
	numbers := []int{8, 9, 99, 28, 739399, 77837}

	for _, n := range numbers {
		expected := false

		actual := checkTwoSidePrime(n)

		if actual != expected {
			t.Errorf("Expected of %v to be %t but got %t", n, expected, actual)
		}
	}
}
