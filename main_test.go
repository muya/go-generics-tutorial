package main

import (
	"testing"
)

var intTests = []struct {
	name     string
	vals     map[string]int64
	expected int64
}{
	{
		name: "all positives",
		vals: map[string]int64{
			"three":       3,
			"fifty-three": 53,
			"five":        5,
		},
		expected: 61,
	},
}

var floatTests = []struct {
	name     string
	vals     map[string]float64
	expected float64
}{
	{
		name: "all positives",
		vals: map[string]float64{
			"three.point.five":                     3.5,
			"three-thirty-two.point.one.two.three": 332.123,
			"five":                                 5,
		},
		expected: 340.623,
	},
}

func Test_SumInts(t *testing.T) {
	for _, test := range intTests {
		t.Run(test.name, func(t *testing.T) {
			actual := SumInts(test.vals)

			if actual != test.expected {
				t.Errorf("Expected sum to be %d, found %d", test.expected, actual)
			}
		})
	}
}

func Test_SumFloats(t *testing.T) {
	for _, test := range floatTests {
		t.Run(test.name, func(t *testing.T) {
			actual := SumFloats(test.vals)

			if actual != test.expected {
				t.Errorf("Expected sum to be %f, found %f", test.expected, actual)
			}
		})
	}
}

func Test_SumIntsOrFloats(t *testing.T) {
	// Test ints
	for _, test := range intTests {
		t.Run(test.name, func(t *testing.T) {
			actual := SumIntsOrFloats[string, int64](test.vals)

			if actual != test.expected {
				t.Errorf("Expected sum to be %d, found %d", test.expected, actual)
			}
		})
	}

	// Test floats
	for _, test := range floatTests {
		t.Run(test.name, func(t *testing.T) {
			actual := SumIntsOrFloats[string, float64](test.vals)

			if actual != test.expected {
				t.Errorf("Expected sum to be %f, found %f", test.expected, actual)
			}
		})
	}
}
