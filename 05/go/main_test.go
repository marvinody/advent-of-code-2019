package main

import "testing"

func TestAddSimple(t *testing.T) {
	insns := []int{1, 0, 0, 3, 99}
	expected := []int{1, 0, 0, 2, 99}

	actual, _ := emulator(insns, 1)
	IntSliceEqualityTest(expected, actual, t)
}

func TestMultiplySimple(t *testing.T) {
	insns := []int{2, 3, 0, 3, 99}
	expected := []int{2, 3, 0, 6, 99}

	actual, _ := emulator(insns, 1)
	IntSliceEqualityTest(expected, actual, t)
}

func TestMultiplySimple2(t *testing.T) {
	insns := []int{2, 4, 4, 5, 99, 0}
	expected := []int{2, 4, 4, 5, 99, 9801}

	actual, _ := emulator(insns, 1)
	IntSliceEqualityTest(expected, actual, t)
}

func TestComplex(t *testing.T) {
	insns := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected := []int{30, 1, 1, 4, 2, 5, 6, 0, 99}

	actual, _ := emulator(insns, 1)
	IntSliceEqualityTest(expected, actual, t)

}

func TestInputOutput(t *testing.T) {
	insns := []int{3, 0, 4, 0, 99}
	input := 4
	output := input
	_, actual := emulator(insns, input)
	if output != actual {
		t.Fatalf("Expected %d, but got %d", output, actual)
	}
}

func TestJumpOnePosition(t *testing.T) {
	insns := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	input := 4
	expected := 1
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestJumpOneImmediate(t *testing.T) {
	insns := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	input := 4
	expected := 1
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestJumpZeroPosition(t *testing.T) {
	insns := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	input := 0
	expected := 0
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestJumpZeroImmediate(t *testing.T) {
	insns := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	input := 0
	expected := 0
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestLessThanPositionZero(t *testing.T) {
	insns := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	input := 0
	expected := 1
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestLessThanPositionOne(t *testing.T) {
	insns := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	input := 9
	expected := 0
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestEqualToPositionOne(t *testing.T) {
	insns := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	input := 8
	expected := 1
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestEqualToImmediateOne(t *testing.T) {
	insns := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	input := 8
	expected := 1
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}
func TestEqualToPositionZero(t *testing.T) {
	insns := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	input := 7
	expected := 0
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestComplex999(t *testing.T) {
	insns := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	input := 7
	expected := 999
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestComplex1000(t *testing.T) {
	insns := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	input := 8
	expected := 1000
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
}

func TestComplex1001(t *testing.T) {
	insns := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	input := 9
	expected := 1001
	_, actual := emulator(insns, input)
	if expected != actual {
		t.Fatalf("Expected %d, but got %d", expected, actual)
	}
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
