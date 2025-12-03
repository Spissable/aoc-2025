package day03

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		var first, second int
		for i, c := range line {
			current, _ := strconv.Atoi(string(c))
			if current > first && i < len(line)-1 {
				first = current
				second = 0
			} else if current > second {
				second = current
			}
		}

		joltage, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))
		result += joltage
	}

	return result
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}
