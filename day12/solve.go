package day12

import (
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) (result int) {
	regions := newRegions(input)
	for _, r := range regions {
		if r.fits() {
			result++
		}
	}
	return result
}

type region struct {
	width  int
	length int
	gifts  []int
}

func newRegions(input string) (result []region) {
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		parts := strings.Split(line, "x")
		if len(parts) == 1 {
			continue
		}

		width, _ := strconv.Atoi(parts[0])
		parts = strings.Split(parts[1], ":")
		length, _ := strconv.Atoi(parts[0])
		parts = strings.Fields(parts[1])
		var gifts []int
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			gifts = append(gifts, num)
		}

		result = append(result, region{
			width:  width,
			length: length,
			gifts:  gifts,
		})
	}
	return result
}

func (r region) fits() bool {
	var giftWidth int
	for _, g := range r.gifts {
		giftWidth += g
	}

	return r.width*r.length >= giftWidth*8
}
