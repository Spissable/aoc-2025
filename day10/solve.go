package day10

import (
	"reflect"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	machines := newMachines(input)
	for _, m := range machines {
		result += getMinButtons(m)
	}
	return result
}

func SolvePuzzle2(input string) int {
	// TODO: solve puzzle 2
	return 0
}

type machine struct {
	lightsTarget []bool
	buttons      [][]int
	joltages     []int
}

func getMinButtons(m machine) int {
	for k := 1; k <= len(m.buttons); k++ {
		for _, c := range combinationsKIter(m.buttons, k) {
			current := make([]bool, len(m.lightsTarget))
			for _, buttons := range c {
				for _, button := range buttons {
					current[button] = !current[button]
				}
			}

			if reflect.DeepEqual(current, m.lightsTarget) {
				return k
			}
		}
	}

	return 0
}

func newMachines(input string) (result []machine) {
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		var (
			lightsTarget []bool
			buttons      [][]int
			joltages     []int
		)
		for i := 0; i < len(line); i++ {
			if line[i] == '[' {
				i++
				for line[i] != ']' {
					lightsTarget = append(lightsTarget, line[i] == '#')
					i++
				}
			}
			if line[i] == '(' {
				i++
				var button []int
				for line[i] != ')' {
					if line[i] == ',' {
						i++
					}
					num, _ := strconv.Atoi(string(line[i]))
					button = append(button, num)
					i++
				}
				buttons = append(buttons, button)
			}
			if line[i] == '{' {
				i++
				var joltageStr string
				for line[i] != '}' {
					joltageStr += string(line[i])
					i++
				}
				parts := strings.SplitSeq(joltageStr, ",")
				for part := range parts {
					num, _ := strconv.Atoi(part)
					joltages = append(joltages, num)
				}
			}
		}
		result = append(result, machine{
			lightsTarget: lightsTarget,
			buttons:      buttons,
			joltages:     joltages,
		})
	}

	return result
}

func combinationsKIter(items [][]int, k int) [][][]int {
	n := len(items)
	if k <= 0 || k > n {
		return nil
	}

	indices := make([]int, k)
	for i := range k {
		indices[i] = i
	}

	var result [][][]int
	for {
		comb := make([][]int, k)
		for i, idx := range indices {
			comb[i] = items[idx]
		}
		result = append(result, comb)

		i := k - 1
		for i >= 0 && indices[i] == i+n-k {
			i--
		}
		if i < 0 {
			break
		}
		indices[i]++
		for j := i + 1; j < k; j++ {
			indices[j] = indices[j-1] + 1
		}
	}
	return result
}
