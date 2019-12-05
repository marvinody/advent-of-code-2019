package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0, 64) // we'll probably have minimum 64 lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	driver(lines)
}

func driver(lines []string) {
	start, _ := strconv.Atoi(lines[0])
	end, _ := strconv.Atoi(lines[1])
	fmt.Println(start, end)
	nums := findNumbers(start, end)
	fmt.Println(nums)
	fmt.Println(len(nums), " numbers")
}

func findNumbers(low, high int) []int {
	var wg sync.WaitGroup
	wg.Add(high - low + 1)
	nums := make([]int, 0, 64)

	for n := low; n <= high; n++ {
		if checkNumber(n) {
			nums = append(nums, n)
		}
	}

	// wg.Wait()
	return nums
}

func checkNumber(n int) bool {
	pairs := pair(numRange((numDigits(n))))
	getDigit := digitGetter(n)
	for _, pair := range pairs {
		// 0-idx is technically the ones digit, and it's on the right
		rightIdx, leftIdx := pair[0], pair[1]
		right, left := getDigit(rightIdx), getDigit(leftIdx)
		if right < left {
			return false
		}
	}
	digitFreq := freq(n)
	for _, count := range digitFreq {
		if count == 2 {
			return true
		}
	}
	return false
}

func digitGetter(n int) func(int) int {
	return func(idx int) int {
		base := int(math.Pow10(idx))
		return (n / base) % 10
	}
}

func numDigits(n int) int {
	return int(math.Floor(math.Log10(float64(n)) + 1))
}

func explode(n int) []int {
	numDigits := numDigits(n)
	digits := make([]int, numDigits)
	getter := digitGetter(n)
	for idx := range digits {
		digits[numDigits-idx-1] = getter(idx)
	}
	return digits
}

func freq(n int) map[int]int {
	freqs := make(map[int]int)
	for _, digit := range explode(n) {
		freqs[digit]++
	}
	return freqs
}

func numRange(n int) []int {
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return nums
}

func pair(arr []int) [][]int {
	pairs := make([][]int, 0, len(arr)-1)
	for i := 1; i < len(arr); i++ {
		pairs = append(pairs, []int{
			i - 1,
			i,
		})
	}
	return pairs
}
