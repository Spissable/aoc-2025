package day02

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	ranges := NewRanges(input)
	for _, r := range ranges {
		for i := r.from; i <= r.to; i++ {
			str := strconv.Itoa(i)
			firstHalf := str[:len(str)/2]
			secondHalf := str[len(str)/2:]
			if firstHalf == secondHalf {
				result += i
			}
		}
	}
	return result
}

func SolvePuzzle2(input string) (result int) {
	ranges := NewRanges(input)
	for _, r := range ranges {
		for currentNum := r.from; currentNum <= r.to; currentNum++ {
			str := strconv.Itoa(currentNum)
			for patternLen := 1; patternLen <= len(str)/2; patternLen++ {
				if len(str)%patternLen != 0 {
					continue
				}

				pattern := str[0:patternLen]
				repeats := true
				for cmpIndex := patternLen; cmpIndex <= len(str)-patternLen; cmpIndex += patternLen {
					if pattern != str[cmpIndex:cmpIndex+patternLen] {
						repeats = false
						break
					}
				}

				if repeats {
					result += currentNum
					break
				}
			}
		}
	}
	return result
}

type rangeNum struct {
	from int
	to   int
}

type ranges []rangeNum

func NewRanges(input string) (result ranges) {
	lines := strings.SplitSeq(input, "\n")

	for line := range lines {
		pairs := strings.SplitSeq(strings.TrimRight(line, ","), ",")
		for pair := range pairs {
			nums := strings.Split(pair, "-")
			from, _ := strconv.Atoi(nums[0])
			to, _ := strconv.Atoi(nums[1])
			result = append(result, rangeNum{from, to})
		}
	}
	return result
}
