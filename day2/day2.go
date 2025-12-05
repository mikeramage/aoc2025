package day2

import (
	"strconv"
	"strings"

	"github.com/mikeramage/aoc2025/utils"
)

type ProductRange struct {
	Start  int
	Finish int
}

func NewProductRange(start, finish int) ProductRange {
	return ProductRange{Start: start, Finish: finish}
}

func testForRepeatedDigits(digits int) bool {

	digitString := strconv.Itoa(digits)

	// Can only have pair of repeated digit sequences if even number of digits
	if len(digitString)%2 != 0 {
		return false
	}

	return digitString[0:len(digitString)/2] == digitString[len(digitString)/2:]

}

func testForMultiRepeatedDigits(digits int) bool {

	digitString := strconv.Itoa(digits)

	for i := 1; i <= len(digitString)/2; i++ {

		// If the current number evenly divides the digit string, it's a possible match
		// for a repeated digit sequence
		if len(digitString)%i == 0 {
			isRepeatedSequence := true
			for j := 0; j <= len(digitString)-2*i; j += i {
				if digitString[j:j+i] != digitString[j+i:j+2*i] {
					isRepeatedSequence = false
					break
				}
			}

			if isRepeatedSequence {
				return true
			}
		}
	}
	return false

}

func Day2() (int, int) {
	lines := utils.Lines("./input/day2.txt")
	if len(lines) != 1 {
		panic("Unexpected number of lines")
	}

	var productRanges []ProductRange
	productRangesList := strings.Split(lines[0], ",")

	for _, productRangesItem := range productRangesList {
		productRange := strings.Split(productRangesItem, "-")
		start, _ := strconv.Atoi(productRange[0])
		finish, _ := strconv.Atoi(productRange[1])
		productRanges = append(productRanges, NewProductRange(start, finish))
	}

	part1 := 0
	part2 := 0

	for _, productRange := range productRanges {
		for i := productRange.Start; i <= productRange.Finish; i++ {
			if testForRepeatedDigits(i) {
				part1 += i
			}

			if testForMultiRepeatedDigits(i) {
				part2 += i
			}

		}
	}

	return part1, part2
}
