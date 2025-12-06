package day3

import (
	"math"

	"github.com/mikeramage/aoc2025/utils"
)

func CalcJoltage(banks []string, numberOfBatteries int) int {
	joltage := 0
	for _, bank := range banks {
		batteries := make([]int, numberOfBatteries)
		for i := 0; i < len(bank)-(numberOfBatteries-1); i++ {
			for j := 0; j < numberOfBatteries; j++ {
				battery := int(bank[i+j] - '0')
				if battery > batteries[j] {
					batteries[j] = battery
					if j != numberOfBatteries-1 {
						batteries[j+1] = 0
					}
				}
			}
		}

		for i := 0; i < numberOfBatteries; i++ {
			joltage += int(math.Pow10(numberOfBatteries-i-1)) * batteries[i]
		}
	}
	return joltage
}

func Day3() (int, int) {
	banks := utils.Lines("./input/day3.txt")

	part1 := CalcJoltage(banks, 2)
	part2 := CalcJoltage(banks, 12)

	return part1, part2
}
