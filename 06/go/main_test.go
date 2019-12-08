package main

import (
	"testing"
)

func TestSimplePath(t *testing.T) {
	lines := []string{
		"COM)B",
		"B)C",
	}

	system := buildSolarSystem(lines)
	lenPathB := system.findOrbitPathLength("B") // B COM
	lenPathC := system.findOrbitPathLength("C") // C B COM
	if lenPathB != 2 {
		t.Errorf("Expected Path length of B to be 2, got %d", lenPathB)
	}
	if lenPathC != 3 {
		t.Errorf("Expected Path length of C to be 3, got %d", lenPathC)
	}

}

func TestLargeSystem(t *testing.T) {
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}
	system := buildSolarSystem(lines)
	total := 0
	for satellite := range system {
		pathLen := system.findOrbitPathLength(satellite)
		total += pathLen - 1
	}
	if total != 42 {
		t.Errorf("Expected total path len to be 42, but got: %d", total)
	}
}

func TestOrbitalMechanics(t *testing.T) {
	lines := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	system := buildSolarSystem(lines)
	pathLen := system.findTransferPathLen("YOU", "SAN")
	if pathLen != 4 {
		t.Errorf("Expected transfer path to be 4, but got %d", pathLen)
	}
}
