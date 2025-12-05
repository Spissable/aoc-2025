package day05

import (
	"sort"
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

func SolvePuzzle2(input string) (result int) {
	ingredients := newIngredients(input)

	sort.Slice(ingredients.freshRanges, func(i, j int) bool {
		return ingredients.freshRanges[i].from < ingredients.freshRanges[j].from
	})

	merged := mergeRanges(ingredients.freshRanges)

	for _, r := range merged {
		result += r.to - r.from + 1
	}

	return result
}

func mergeRanges(ranges []idRange) []idRange {
	if len(ranges) == 0 {
		return nil
	}

	merged := []idRange{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		current := ranges[i]
		last := &merged[len(merged)-1]

		if current.from <= last.to+1 {
			if current.to > last.to {
				last.to = current.to
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
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
