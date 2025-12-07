package day4

import (
	"github.com/mikeramage/aoc2025/utils"
)

type GridContents byte

var ROLL GridContents = '@'
var EMPTY GridContents = '.'

func CountAdjacentRolls(grid [][]GridContents, i, j int) int {
	numAdjacentRolls := 0

	for k := i - 1; k < i+2; k++ {
		if k < 0 || k > len(grid)-1 {
			continue
		}
		for l := j - 1; l < j+2; l++ {
			if l < 0 || l > len(grid[0])-1 || (k == i && l == j) {
				continue
			}

			if grid[k][l] == ROLL {
				numAdjacentRolls++
			}
		}
	}

	return numAdjacentRolls
}

func RemoveRolls(grid [][]GridContents) int {
	numRemoved := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == ROLL && CountAdjacentRolls(grid, i, j) < 4 {
				grid[i][j] = EMPTY
				numRemoved++
			}
		}
	}

	return numRemoved
}

func Day4() (int, int) {
	lines := utils.Lines("./input/day4.txt")
	grid := make([][]GridContents, 0)

	for _, rowContents := range lines {
		row := make([]GridContents, 0)
		for _, item := range rowContents {
			row = append(row, GridContents(item))
		}
		grid = append(grid, row)
	}

	part1 := RemoveRolls(grid)
	part2 := part1
	var previous int

	for part2 != previous {
		previous = part2
		part2 += RemoveRolls(grid)
	}

	return part1, part2
}
