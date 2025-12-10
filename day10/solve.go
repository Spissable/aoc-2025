package day10

import (
	"math"
	"reflect"
	"slices"
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

func SolvePuzzle2(input string) (result int) {
	machines := newMachines(input)
	for _, m := range machines {
		result += solveLPMachine(m)
	}
	return result
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

const (
	inf = 1e18
	eps = 1e-9
)

func simplex(lhs [][]float64, c []float64) (float64, []float64) {
	m := len(lhs)
	n := len(lhs[0]) - 1

	nIndices := make([]int, n+1)
	for i := range n {
		nIndices[i] = i
	}
	nIndices[n] = -1

	bIndices := make([]int, m)
	for i := range m {
		bIndices[i] = n + i
	}

	d := make([][]float64, m+2)
	for i := range d {
		d[i] = make([]float64, n+2)
	}

	for i := range m {
		copy(d[i], lhs[i][:n+1])
		d[i][n+1] = -1.0
	}

	for i := range m {
		d[i][n], d[i][n+1] = d[i][n+1], d[i][n]
	}

	copy(d[m], c[:n])
	d[m+1][n] = 1.0

	pivot := func(r, s int) {
		k := 1.0 / d[r][s]

		for i := 0; i < m+2; i++ {
			if i == r {
				continue
			}
			for j := 0; j < n+2; j++ {
				if j != s {
					d[i][j] -= d[r][j] * d[i][s] * k
				}
			}
		}

		for i := 0; i < n+2; i++ {
			d[r][i] *= k
		}
		for i := 0; i < m+2; i++ {
			d[i][s] *= -k
		}
		d[r][s] = k

		bIndices[r], nIndices[s] = nIndices[s], bIndices[r]
	}

	find := func(pIdx int) bool {
		for {
			bestS := -1
			bestVal := inf
			bestNIdx := int(1e9)

			for i := 0; i <= n; i++ {
				if pIdx != 0 || nIndices[i] != -1 {
					val := d[m+pIdx][i]
					if bestS == -1 || val < bestVal-eps || (math.Abs(val-bestVal) <= eps && nIndices[i] < bestNIdx) {
						bestS = i
						bestVal = val
						bestNIdx = nIndices[i]
					}
				}
			}

			s := bestS
			if d[m+pIdx][s] > -eps {
				return true
			}

			bestR := -1
			bestRatio := inf
			bestBIdx := int(1e9)

			for i := range m {
				if d[i][s] > eps {
					ratio := d[i][n+1] / d[i][s]
					if bestR == -1 || ratio < bestRatio-eps || (math.Abs(ratio-bestRatio) <= eps && bIndices[i] < bestBIdx) {
						bestR = i
						bestRatio = ratio
						bestBIdx = bIndices[i]
					}
				}
			}

			r := bestR
			if r == -1 {
				return false
			}

			pivot(r, s)
		}
	}

	splitR := 0
	minVal := d[0][n+1]
	for i := 1; i < m; i++ {
		if d[i][n+1] < minVal {
			minVal = d[i][n+1]
			splitR = i
		}
	}

	if d[splitR][n+1] < -eps {
		pivot(splitR, n)
		if !find(1) || d[m+1][n+1] < -eps {
			return -inf, nil
		}

		for i := range m {
			if bIndices[i] == -1 {
				bestS := 0
				bestVal := d[i][0]
				bestNIdx := nIndices[0]

				for j := 1; j < n; j++ {
					if d[i][j] < bestVal-eps || (math.Abs(d[i][j]-bestVal) <= eps && nIndices[j] < bestNIdx) {
						bestS = j
						bestVal = d[i][j]
						bestNIdx = nIndices[j]
					}
				}
				pivot(i, bestS)
			}
		}
	}

	if find(0) {
		x := make([]float64, n)
		for i := range m {
			if bIndices[i] >= 0 && bIndices[i] < n {
				x[bIndices[i]] = d[i][n+1]
			}
		}

		sumVal := 0.0
		for i := range n {
			sumVal += c[i] * x[i]
		}
		return sumVal, x
	}

	return -inf, nil
}

func solveILPBnB(initialA [][]float64, objCoeffs []float64) int {
	bestVal := inf
	stack := [][][]float64{initialA}

	for len(stack) > 0 {
		currentA := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		val, xOpt := simplex(currentA, objCoeffs)

		if val == -inf || val >= bestVal-eps {
			continue
		}

		fractionalIdx := -1
		fractionalVal := 0.0

		if xOpt != nil {
			for i, xv := range xOpt {
				if math.Abs(xv-math.Round(xv)) > eps {
					fractionalIdx = i
					fractionalVal = xv
					break
				}
			}

			if fractionalIdx != -1 {
				floorV := math.Floor(fractionalVal)
				nCols := len(currentA[0])

				row1 := make([]float64, nCols)
				row1[fractionalIdx] = 1.0
				row1[nCols-1] = floorV
				a1 := make([][]float64, len(currentA)+1)
				for i := range currentA {
					a1[i] = make([]float64, nCols)
					copy(a1[i], currentA[i])
				}
				a1[len(currentA)] = row1
				stack = append(stack, a1)

				ceilV := math.Ceil(fractionalVal)
				row2 := make([]float64, nCols)
				row2[fractionalIdx] = -1.0
				row2[nCols-1] = -ceilV
				a2 := make([][]float64, len(currentA)+1)
				for i := range currentA {
					a2[i] = make([]float64, nCols)
					copy(a2[i], currentA[i])
				}
				a2[len(currentA)] = row2
				stack = append(stack, a2)
			} else if val < bestVal {
				bestVal = val
			}
		}
	}

	if bestVal == inf {
		return 0
	}
	return int(math.Round(bestVal))
}

func solveLPMachine(m machine) int {
	numGoals := len(m.joltages)
	numButtons := len(m.buttons)

	rows := 2*numGoals + numButtons
	cols := numButtons + 1

	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}

	for j := range numButtons {
		rowIdx := rows - 1 - j
		matrix[rowIdx][j] = -1.0
	}

	for j, buttons := range m.buttons {
		for i := range numGoals {
			if slices.Contains(buttons, i) {
				matrix[i][j] = 1.0
				matrix[i+numGoals][j] = -1.0
			}
		}
	}

	for i := range numGoals {
		val := float64(m.joltages[i])
		matrix[i][cols-1] = val
		matrix[i+numGoals][cols-1] = -val
	}

	objCoeffs := make([]float64, numButtons)
	for i := range objCoeffs {
		objCoeffs[i] = 1.0
	}

	return solveILPBnB(matrix, objCoeffs)
}
