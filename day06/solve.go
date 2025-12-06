package day06

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	c := newCalc(input)
	for _, calc := range c {
		subTotal := calc.nums[0]
		for _, num := range calc.nums[1:] {
			if calc.add {
				subTotal += num
			} else {
				subTotal *= num
			}
		}

		result += subTotal
	}

	return result
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}

type eval struct {
	nums []int
	add  bool
}

func newCalc(input string) (result []eval) {
	lines := strings.Split(input, "\n")

	for lineI, line := range lines {
		elems := strings.Fields(line)
		for elemI, elem := range elems {
			if lineI == 0 {
				result = append(result, eval{})
			}

			num, err := strconv.Atoi(strings.TrimSpace(elem))
			if err != nil {
				result[elemI].add = elem == "+"
				continue
			}

			result[elemI].nums = append(result[elemI].nums, num)
		}
	}

	return result
}
