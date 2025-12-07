/*
Copyright © 2025 Mike Ramage <mike.ramage@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/mikeramage/aoc2025/day1"
	// "github.com/mikeramage/aoc2025/day10"
	// "github.com/mikeramage/aoc2025/day11"
	// "github.com/mikeramage/aoc2025/day12"
	"github.com/mikeramage/aoc2025/day2"
	"github.com/mikeramage/aoc2025/day3"

	"github.com/mikeramage/aoc2025/day4"
	// "github.com/mikeramage/aoc2025/day5"
	// "github.com/mikeramage/aoc2025/day6"
	// "github.com/mikeramage/aoc2025/day7"
	// "github.com/mikeramage/aoc2025/day8"
	// "github.com/mikeramage/aoc2025/day9"
	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve Advent of Code",
	Long: `Solve Advent of Code for the specified day(s) or range of days, outputs 
the associated solutions and visualizations`,
	Run: func(cmd *cobra.Command, args []string) {
		solutions := []func() (int, int){
			day1.Day1,
			day2.Day2,
			day3.Day3,
			day4.Day4,
			// day5.Day5,
			// day6.Day6,
			// day7.Day7,
			// day8.Day8,
			// day9.Day9,
			// day10.Day10,
			// day11.Day11,
			// day12.Day12,
		}

		if day == -1 {
			totalStartTime := time.Now()
			for i, solution := range solutions {
				doDay(i+1, solution)
			}
			totalEndTime := time.Now()
			totalElapsed := totalEndTime.Sub(totalStartTime)
			fmt.Println("All days took", totalElapsed)
		} else {
			doDay(day, solutions[day-1])
		}
	},
}

func doDay(day int, solution func() (int, int)) {
	startTime := time.Now()
	part1, part2 := solution()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("Day%v took %v\n", day, elapsed)
	fmt.Println("  Part 1:", part1)
	fmt.Println("  Part 2:", part2)
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(solveCmd)
	solveCmd.Flags().IntVarP(&day, "day", "d", -1, "Solve the specified day only")
}
