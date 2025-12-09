package day09

import (
	"sort"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	pairs := newPairs(newCoords(input))
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].area > pairs[b].area
	})

	return pairs[0].area
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}

type coord struct {
	x, y int
}

type pair struct {
	a, b coord
	area int
}

func newCoords(input string) (result []coord) {
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		fields := strings.Split(line, ",")
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])

		result = append(result, coord{x: x, y: y})
	}

	return result
}

func getArea(a, b coord) int {
	x := a.x - b.x + 1
	if x < 0 {
		x *= -1
	}

	y := a.y - b.y + 1
	if y < 0 {
		y *= -1
	}

	return x * y
}

func newPairs(points []coord) (result []pair) {
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			result = append(result, pair{
				a:    a,
				b:    b,
				area: getArea(a, b),
			})
		}
	}

	return result
}
