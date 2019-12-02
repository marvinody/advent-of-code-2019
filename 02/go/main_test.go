package main

import "testing"

func TestAddSimple(t *testing.T) {
	insns := []int{1, 0, 0, 3, 99}
	expected := []int{1, 0, 0, 2, 99}

	actual := emulator(insns)
	IntSliceEqualityTest(expected, actual, t)
}

func TestMultiplySimple(t *testing.T) {
	insns := []int{2, 3, 0, 3, 99}
	expected := []int{2, 3, 0, 6, 99}

	actual := emulator(insns)
	IntSliceEqualityTest(expected, actual, t)
}

func TestMultiplySimple2(t *testing.T) {
	insns := []int{2, 4, 4, 5, 99, 0}
	expected := []int{2, 4, 4, 5, 99, 9801}

	actual := emulator(insns)
	IntSliceEqualityTest(expected, actual, t)
}

func TestComplex(t *testing.T) {
	insns := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	actual := emulator(insns)
	IntSliceEqualityTest(expected, actual, t)

}

func IntSliceEqualityTest(a, b []int, t *testing.T) {
	if len(a) != len(b) {
		t.Fatalf("Different lengths: %d <> %d", len(a), len(b))
	}

	for idx := range a {
		if a[idx] != b[idx] {
			t.Fatalf("Arrays differ at index: %d, %d <> %d", idx, a[idx], b[idx])
		}
	}
}
