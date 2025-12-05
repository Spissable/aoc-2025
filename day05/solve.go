package day05

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	ingredients := newIngredients(input)
	for _, ingredient := range ingredients.available {
		if ingredients.isFresh(ingredient) {
			result++
		}
	}
	return result
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}

type idRange struct {
	from, to int
}

type ingredients struct {
	freshRanges []idRange
	available   []int
}

func newIngredients(input string) (result ingredients) {
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			from, _ := strconv.Atoi(parts[0])
			to, _ := strconv.Atoi(parts[1])
			result.freshRanges = append(result.freshRanges, idRange{from: from, to: to})
		} else {
			id, _ := strconv.Atoi(parts[0])
			result.available = append(result.available, id)
		}
	}
	return result
}

func (i ingredients) isFresh(id int) bool {
	for _, r := range i.freshRanges {
		if id >= r.from && id <= r.to {
			return true
		}
	}

	return false
}
