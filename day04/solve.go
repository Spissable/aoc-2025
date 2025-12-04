package day04

import (
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	p := newPapers(input)

	for coords, _ := range p {
		if len(p.getNeighbors(coords.x, coords.y)) < 4 {
			result++
		}
	}

	return result
}

func SolvePuzzle2(input string) (result int) {
	p := newPapers(input)
	didRemove := true

	for didRemove == true {
		didRemove = false

		for coords, _ := range p {
			if len(p.getNeighbors(coords.x, coords.y)) < 4 {
				result++
				delete(p, coords)
				didRemove = true
			}
		}
	}

	return result
}

type coord struct{ x, y int }
type papers map[coord]bool

func newPapers(input string) papers {
	result := make(map[coord]bool)
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, c := range line {
			if c == '@' {
				result[coord{x: x, y: y}] = true
			}
		}
	}

	return result
}

func (p papers) getNeighbors(x, y int) (result []coord) {
	for _, c := range neighborCombinations {
		neighbor := coord{x: x + c.x, y: y + c.y}
		if p[neighbor] {
			result = append(result, neighbor)
		}
	}

	return result
}

var neighborCombinations = []coord{
	{
		x: 0,
		y: 1,
	},
	{
		x: 1,
		y: 0,
	},
	{
		x: 1,
		y: 1,
	},
	{
		x: 0,
		y: -1,
	},
	{
		x: -1,
		y: 0,
	},
	{
		x: -1,
		y: -1,
	},
	{
		x: -1,
		y: 1,
	},
	{
		x: 1,
		y: -1,
	},
}
