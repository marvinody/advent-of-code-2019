package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 64) // we'll probably have minimum 64 lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	driver(lines)

}

// the body of this function will change all the time!
func driver(lines []string) {

	insns := make([]int, 0, len(lines))
	nums := strings.Split(lines[0], ",")
	for _, num := range nums {
		insn, _ := strconv.Atoi(num)
		insns = append(insns, insn)
	}

	input := 5
	_, output := emulator(insns, input)
	fmt.Println("p1 output:", output)

}

// op codes
const (
	ADD           int = 1
	MULTIPLY          = 2
	WRITE             = 3 // writes input into int array
	READ              = 4 // reads some int and will output it
	JUMP_IF_TRUE      = 5
	JUMP_IF_FALSE     = 6
	LESS_THAN         = 7
	EQUAL_TO          = 8
	HALT              = 99
)

// addressing modes
const (
	POSITION  int = 0
	IMMEDIATE int = 1
)

var opLength = map[int]int{
	ADD:           4,
	MULTIPLY:      4,
	READ:          2,
	WRITE:         2,
	HALT:          1,
	JUMP_IF_TRUE:  3,
	JUMP_IF_FALSE: 3,
	LESS_THAN:     4,
	EQUAL_TO:      4,
}

func emulator(insns []int, input int) ([]int, int) {
	PC := 0 // where we at in the program
	output := 0
	for {
		insn := insns[PC]
		op := insn % 100
		getMode := digitGetter(insn)
		// exit quickly to prevent out of bounds from next statements
		if insn == HALT {
			return insns, output
		}

		// the next few 'bytes' are our addresses/immediates
		leftVal, rightVal, outAddress := 0, 0, 0
		leftVal = insns[PC+1]
		if false ||
			op == ADD ||
			op == MULTIPLY ||
			op == JUMP_IF_TRUE ||
			op == JUMP_IF_FALSE ||
			op == LESS_THAN ||
			op == EQUAL_TO {
			rightVal = insns[PC+2]
			outAddress = insns[PC+3]    // this will probably always be position
			if getMode(2) == POSITION { // for digit in the hundreds (10 ^ 2)
				leftVal = insns[leftVal] // access the value stored at that address
			}
			if getMode(3) == POSITION { // for digit in the thousands (10 ^ 3)
				rightVal = insns[rightVal]
			}
		}

		if op == READ {
			if getMode(2) == POSITION { // for digit in the hundreds (10 ^ 2)
				leftVal = insns[leftVal] // access the value stored at that address
			}
		}

		switch op {
		case ADD:
			result := leftVal + rightVal
			insns[outAddress] = result
		case MULTIPLY:
			result := leftVal * rightVal
			insns[outAddress] = result
		case WRITE:
			insns[leftVal] = input
		case READ:
			output = leftVal
		case JUMP_IF_TRUE:
			if leftVal != 0 {
				PC = rightVal
				continue
			}
		case JUMP_IF_FALSE:
			if leftVal == 0 {
				PC = rightVal
				continue
			}
		case LESS_THAN:
			insns[outAddress] = b2i(leftVal < rightVal)
		case EQUAL_TO:
			insns[outAddress] = b2i(leftVal == rightVal)
		}

		// each op is 4 'bytes' long so we skip the section each loop
		PC += opLength[op]
	}
}

// from day 4
func digitGetter(n int) func(int) int {
	return func(idx int) int {
		base := int(math.Pow10(idx))
		return (n / base) % 10
	}
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
