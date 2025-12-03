package day1

import (
	"strconv"
	"strings"

	"github.com/mikeramage/aoc2025/utils"
)

type Rotation struct {
	Clockwise bool
	Distance  int
}

func NewRotation(clockwise bool, distance int) Rotation {
	return Rotation{Clockwise: clockwise, Distance: distance}
}

func ApplyRotation(oldValue int, rotation Rotation) (newValue int, numZeroClicks int) {
	numZeroClicks = 0

	//First rotate to zero if not already at 0 and have enough distance to reach 0
	if oldValue != 0 {
		if rotation.Clockwise && (rotation.Distance >= (100 - oldValue)) {
			rotation.Distance -= (100 - oldValue)
			numZeroClicks++
			oldValue = 0
		} else if !rotation.Clockwise && (rotation.Distance >= oldValue) {
			rotation.Distance -= oldValue
			numZeroClicks++
			oldValue = 0
		}
	}

	numZeroClicks += (int)(rotation.Distance / 100)

	if rotation.Clockwise {
		newValue = (oldValue + rotation.Distance) % 100
	} else {
		newValue = (oldValue - rotation.Distance) % 100
	}

	if newValue < 0 {
		newValue += 100
	}

	return newValue, numZeroClicks
}

func Day1() (int, int) {
	lines := utils.Lines("./input/day1.txt")
	var rotations []Rotation

	for _, line := range lines {
		input := strings.TrimSpace(line)
		clockwise := input[0] == 'R'
		distance, _ := strconv.Atoi(input[1:])
		rotations = append(rotations, NewRotation(clockwise, distance))
	}

	part1 := 0
	part2 := 0

	dialValue := 50
	for i := 0; i < len(rotations); i++ {
		var numZeroClicks int
		dialValue, numZeroClicks = ApplyRotation(dialValue, rotations[i])

		if dialValue == 0 {
			part1++
		}

		part2 += numZeroClicks
	}

	return part1, part2
}
