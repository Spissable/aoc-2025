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
	coords := newCoords(input)
	pairs := newPairs(coords)
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].area > pairs[b].area
	})
	edges := make([]coord, 0)
	edges = append(edges, coords...)
	coords = append(coords, coords[0])
	for i, p := range coords[:len(coords)-1] {
		edges = append(edges, getEdges(p, coords[i+1])...)
	}

	for j := range pairs {
		if isValid(pairs[j], edges) {
			return pairs[j].area
		}
	}

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
	return (max(a.x, b.x) - min(a.x, b.x) + 1) * (max(a.y, b.y) - min(a.y, b.y) + 1)
}

func newPairs(coords []coord) (result []pair) {
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			a := coords[i]
			b := coords[j]
			result = append(result, pair{
				a:    a,
				b:    b,
				area: getArea(a, b),
			})
		}
	}

	return result
}

func i(p pair, e coord) bool {
	return e.x < max(p.a.x, p.b.x) && e.x > min(p.a.x, p.b.x) && e.y < max(p.a.y, p.b.y) && e.y > min(p.a.y, p.b.y)
}

func getEdges(a, b coord) (c []coord) {
	if a.x == b.x {
		for i := min(a.y, b.y) + 1; i < max(a.y, b.y); i++ {
			c = append(c, coord{a.x, i})
		}
		return c
	}
	for i := min(a.x, b.x) + 1; i < max(a.x, b.x); i++ {
		c = append(c, coord{i, a.y})
	}
	return c
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValid(p pair, edges []coord) bool {
	for _, e := range edges {
		if e.x < max(p.a.x, p.b.x) && e.x > min(p.a.x, p.b.x) && e.y < max(p.a.y, p.b.y) && e.y > min(p.a.y, p.b.y) {
			return false
		}
	}

	return true
}
