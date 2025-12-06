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

func SolvePuzzle2(input string) (result int) {
	c := newCalc2(input)
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

func newCalc2(input string) (result []eval) {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return result
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	columnRanges := [][]int{}

	inColumn := false
	start := 0

	for pos := 0; pos < maxLen; pos++ {
		hasContent := false
		for _, line := range lines {
			if pos < len(line) && line[pos] != ' ' {
				hasContent = true
				break
			}
		}

		if hasContent && !inColumn {
			start = pos
			inColumn = true
		} else if !hasContent && inColumn {
			columnRanges = append(columnRanges, []int{start, pos})
			inColumn = false
		}
	}

	if inColumn {
		columnRanges = append(columnRanges, []int{start, maxLen})
	}

	for colIdx := len(columnRanges) - 1; colIdx >= 0; colIdx-- {
		colRange := columnRanges[colIdx]
		startPos := colRange[0]
		endPos := colRange[1]

		eval := eval{}

		for pos := endPos - 1; pos >= startPos; pos-- {
			var digits []string
			var operator string

			for _, line := range lines {
				if pos < len(line) && line[pos] != ' ' {
					char := string(line[pos])
					if char == "+" || char == "*" {
						operator = char
					} else {
						digits = append(digits, char)
					}
				}
			}

			if len(digits) > 0 {
				numberStr := strings.Join(digits, "")
				if num, err := strconv.Atoi(numberStr); err == nil {
					eval.nums = append(eval.nums, num)
				}
			}

			if operator != "" {
				eval.add = operator == "+"
			}
		}

		if len(eval.nums) > 0 {
			result = append(result, eval)
		}
	}

	return result
}
