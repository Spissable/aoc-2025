package day11

import (
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	m := newMachines(input)
	return m.countPaths("you")
}

func SolvePuzzle2(input string) (result int) {
	m := newMachines(input)
	return m.countPaths2("svr", false, false, make(map[cacheKey]int))
}

type cacheKey struct {
	name string
	dac  bool
	fft  bool
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

func (m machines) countPaths2(name string, dac, fft bool, cache map[cacheKey]int) int {
	key := cacheKey{name, dac, fft}
	if cached, ok := cache[key]; ok {
		return cached
	}

	current := m[name]
	if current[0] == "out" {
		if dac && fft {
			cache[key] = 1
			return 1
		}

		cache[key] = 0
		return 0
	}

	var result int
	for _, c := range current {
		newDac := dac || c == "dac"
		newFft := fft || c == "fft"
		result += m.countPaths2(c, newDac, newFft, cache)
	}

	cache[key] = result
	return result
}
