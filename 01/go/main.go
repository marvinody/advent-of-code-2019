package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	totalFuel := 0
	for _, str := range lines {
		mass, _ := strconv.Atoi(str)
		fuel := findFuelForMass(mass)
		totalFuel += fuel
	}
	fmt.Println("TotalFuel:", totalFuel)

}

func findFuelForMass(mass int) int {
	fuelNeededForMass := mass/3 - 2
	if fuelNeededForMass <= 6 {
		return fuelNeededForMass
	}
	return fuelNeededForMass + findFuelForMass(fuelNeededForMass)
}
