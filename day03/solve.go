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

func SolvePuzzle2(input string) (result int) {
	lines := strings.SplitSeq(input, "\n")
	length := 12

	for line := range lines {
		batteries := make([]int, length)
		for charIndex, c := range line {
			current, _ := strconv.Atoi(string(c))
			reset := false
			for batteryIndex, battery := range batteries {
				if reset {
					if battery == 0 {
						break
					}
					batteries[batteryIndex] = 0
					continue
				}

				if battery == 0 {
					batteries[batteryIndex] = current
					break
				}

				if current > battery {
					if charIndex < len(line)-length+batteryIndex+1 {
						batteries[batteryIndex] = current
						reset = true
					} else if batteryIndex == length-1 {
						batteries[batteryIndex] = current
					}
				}
			}
		}

		var joltageStr string
		for _, battery := range batteries {
			joltageStr += strconv.Itoa(battery)
		}

		joltage, _ := strconv.Atoi(joltageStr)
		result += joltage
	}

	return result
}
