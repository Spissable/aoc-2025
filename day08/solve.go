package day08

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	pairs := newPairs(newBoxes(input))

	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].distance < pairs[b].distance
	})

	var circuits circuits
	for _, pair := range pairs[:1000] {
		circuits.addPair(pair)
	}

	sort.Slice(circuits, func(a, b int) bool {
		return len(circuits[a]) > len(circuits[b])
	})

	result = 1
	for _, c := range circuits[:3] {
		result *= len(c)
	}

	return result
}

func SolvePuzzle2(input string) int {
	boxes := newBoxes(input)
	pairs := newPairs(boxes)

	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].distance < pairs[b].distance
	})

	var circuits circuits
	for _, pair := range pairs {
		circuits.addPair(pair)

		if len(circuits) == 1 && len(circuits[0]) == len(boxes) {
			return pair.p.x * pair.q.x
		}
	}

	return 0
}

type coord struct {
	x, y, z int
}

func (c coord) equal(c2 coord) bool {
	return c.x == c2.x && c.y == c2.y && c.z == c2.z
}

func newBoxes(input string) (result []coord) {
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		fields := strings.Split(line, ",")
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		z, _ := strconv.Atoi(fields[2])

		result = append(result, coord{x: x, y: y, z: z})
	}

	return result
}

func getDistance(p, q coord) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2) + math.Pow(float64(p.z-q.z), 2))
}

type pair struct {
	p, q     coord
	distance float64
}

func newPairs(boxes []coord) (result []pair) {
	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			p := boxes[i]
			q := boxes[j]
			result = append(result, pair{
				p:        p,
				q:        q,
				distance: getDistance(p, q),
			})
		}
	}

	return result
}

type circuits [][]coord

func (c *circuits) addPair(pair pair) {
	var pIndex, qIndex int = -1, -1

	for i, circuit := range *c {
		for _, coord := range circuit {
			if coord.equal(pair.p) {
				pIndex = i
			}
			if coord.equal(pair.q) {
				qIndex = i
			}
		}
	}

	// Both are in the same circuit
	if pIndex != -1 && pIndex == qIndex {
		return
	}

	// Both are in different circuits - merge them
	if pIndex != -1 && qIndex != -1 {
		(*c)[pIndex] = append((*c)[pIndex], (*c)[qIndex]...)
		*c = append((*c)[:qIndex], (*c)[qIndex+1:]...)
		return
	}

	// Only p is in a circuit - add q to it
	if pIndex != -1 {
		(*c)[pIndex] = append((*c)[pIndex], pair.q)
		return
	}

	// Only q is in a circuit - add p to it
	if qIndex != -1 {
		(*c)[qIndex] = append((*c)[qIndex], pair.p)
		return
	}

	// Neither is in a circuit - create new circuit
	*c = append(*c, []coord{pair.p, pair.q})
}
