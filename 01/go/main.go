package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 64) // we'll probably have minimum 64 lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(lines)
}
