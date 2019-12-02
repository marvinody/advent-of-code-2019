package main

import (
	"bufio"
	"fmt"
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
	output := emulator(insns)
	fmt.Println("output:", output)

}

// op codes
const (
	ADD      int = 1
	MULTIPLY     = 2
	HALT         = 99
)

func emulator(insns []int) []int {
	PC := 0 // where we at in the program
	for {
		insn := insns[PC]
		// exit quickly to prevent out of bounds from next statements
		if insn == HALT {
			return insns
		}

		// the next few 'bytes' are our addresses
		leftAddr := insns[PC+1]
		rightAddr := insns[PC+2]
		outAddress := insns[PC+3]
		switch insn {
		case ADD:
			result := insns[leftAddr] + insns[rightAddr]
			insns[outAddress] = result
		case MULTIPLY:
			result := insns[leftAddr] * insns[rightAddr]
			insns[outAddress] = result
		}

		// each op is 4 'bytes' long so we skip the section each loop
		PC += 4

	}
}
