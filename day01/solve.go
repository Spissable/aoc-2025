package day01

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) int {
	rotations := newRotations(input)
	pos := 50
	result := 0

	for _, r := range rotations {
		switch r.direction {
		case 'L':
			pos -= r.distance
		case 'R':
			pos += r.distance
		}

		if pos%100 == 0 {
			result++
		}
	}

	return result
}

func SolvePuzzle2(input string) int {
	rotations := newRotations(input)
	pos := 50
	result := 0

	for _, r := range rotations {
		switch r.direction {
		case 'L':
			for i := 0; i < r.distance; i++ {
				pos--
				if pos%100 == 0 {
					result++
				}
			}
		case 'R':
			for i := 0; i < r.distance; i++ {
				pos++
				if pos%100 == 0 {
					result++
				}
			}
		}
	}

	return result
}

type rotation struct {
	direction rune
	distance  int
}

func newRotations(input string) (result []rotation) {
	lines := strings.SplitSeq(input, "\n")
	for l := range lines {
		distance, err := strconv.Atoi(l[1:])
		if err != nil {
			panic(err)
		}

		switch l[0] {
		case 'L':
			result = append(result, rotation{'L', distance})
		case 'R':
			result = append(result, rotation{'R', distance})
		}
	}

	return result
}
