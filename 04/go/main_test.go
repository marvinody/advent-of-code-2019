package main

import (
	"testing"
)

type testCase struct {
	n        int
	expected bool
}

func TestSomeNums(t *testing.T) {
	cases := []testCase{
		testCase{111111, true},
		testCase{223450, false},
		testCase{123479, false},
		testCase{122345, true},
		testCase{111123, true},
		testCase{135679, false},
	}

	for _, testCase := range cases {
		actual := checkNumber(testCase.n)
		if actual != testCase.expected {
			t.Errorf("Expected %d to equal %t", testCase.n, testCase.expected)
		}
	}

}
