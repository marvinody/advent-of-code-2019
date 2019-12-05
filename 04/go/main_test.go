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
		testCase{111111, false},
		testCase{223450, false},
		testCase{123479, false},
		testCase{122345, true},
		testCase{112233, true},
		testCase{111122, true},
		testCase{111123, false},
		testCase{135679, false},
		testCase{123444, false},
	}

	for _, testCase := range cases {
		actual := checkNumber(testCase.n)
		if actual != testCase.expected {
			t.Errorf("Expected %d to equal %t", testCase.n, testCase.expected)
		}
	}

}

func TestExplode(t *testing.T) {
	n := 13579
	digits := []int{1, 3, 5, 7, 9}

	for idx, digit := range explode(n) {
		if digit != digits[idx] {
			t.Errorf("Error in explode")
		}
	}
}
