package day11

import (
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	m := newMachines(input)
	return m.countPaths("you")
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}

type machines map[string][]string

func newMachines(input string) machines {
	m := make(machines)

	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		parts := strings.Split(line, ":")
		name := parts[0]

		outputs := strings.Split(strings.TrimSpace(parts[1]), " ")
		m[name] = outputs
	}

	return m
}

func (m machines) countPaths(name string) int {
	current := m[name]
	if current[0] == "out" {
		return 1
	}

	var result int
	for _, c := range current {
		result += m.countPaths(c)
	}

	return result
}
