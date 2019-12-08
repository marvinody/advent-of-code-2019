package main

import (
	"bufio"
	"fmt"
	"os"
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
	system := buildSolarSystem(lines)
	total := 0
	for satellite := range system {
		pathLen := system.findOrbitPathLength(satellite)
		total += pathLen - 1
	}

	fmt.Println("p1: ", total)
	fmt.Println("p2:", system.findTransferPathLen("YOU", "SAN"))
}

// moon -> planet
type solarSystem map[string]string

func (system solarSystem) findTransferPathLen(start, end string) int {
	startPath := reversed(system.findOrbitPath(start)[1:])
	endPath := reversed(system.findOrbitPath(end)[1:])
	idx := 0
	for startPath[idx] == endPath[idx] {
		idx++
	}
	return (len(startPath) - idx) + (len(endPath) - idx)

}

func (system solarSystem) findOrbitPath(start string) []string {
	curr := start
	path := []string{}
	for curr != "" {
		path = append(path, curr)
		curr = system[curr]
	}
	return path
}

func (system solarSystem) findOrbitPathLength(start string) int {
	return len(system.findOrbitPath(start))
}

func buildSolarSystem(lines []string) solarSystem {
	system := make(map[string]string)
	for _, line := range lines {
		stuff := strings.Split(line, ")")
		planet, moon := stuff[0], stuff[1]
		system[moon] = planet
	}
	return system
}

func reversed(s []string) []string {
	r := make([]string, len(s))
	l := len(s)
	for idx := range s {
		r[l-idx-1] = s[idx]
	}
	return r
}
