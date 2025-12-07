package day07

import (
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	d := newDiagram(input)
	return d.walk(coord{x: d.start.x, y: d.start.y + 1}, make(map[coord]bool))
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}

type coord struct {
	x, y int
}

type diagram struct {
	start         coord
	splitters     map[coord]bool
	width, height int
}

func newDiagram(input string) diagram {
	lines := strings.Split(input, "\n")
	result := diagram{
		splitters: make(map[coord]bool),
		width:     len(lines[0]),
		height:    len(lines),
	}
	for y, line := range lines {
		for x, char := range line {
			switch char {
			case 'S':
				result.start = coord{x: x, y: y}
			case '^':
				result.splitters[coord{x: x, y: y}] = true
			}
		}
	}
	return result
}

func (d diagram) walk(pos coord, walked map[coord]bool) int {
	if walked[pos] || pos.x > d.width-1 || pos.x < 0 || pos.y > d.height-1 {
		return 0
	}

	walked[pos] = true

	if d.splitters[pos] {
		return 1 + d.walk(coord{x: pos.x - 1, y: pos.y + 1}, walked) + d.walk(coord{x: pos.x + 1, y: pos.y + 1}, walked)
	}

	return d.walk(coord{x: pos.x, y: pos.y + 1}, walked)
}
