package main

import (
	"fmt"

	"github.com/t0nyandre/adventofcode/utils"
)

func part1(input []string) int {
	var sum int
	grid := make([][]rune, len(input))

	word := "XMAS"

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{-1, -1},
		{1, -1},
		{1, 1},
	}

	for k, line := range input {
		grid[k] = []rune(line)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, direction := range directions {
				ri := direction[0]
				ci := direction[1]
				xmas := true

				for chi := 0; chi < len(word); chi++ {
					rOffset := row + (ri * chi)
					cOffset := col + (ci * chi)

					if rOffset < 0 || rOffset >= len(grid) || cOffset < 0 || cOffset >= len(grid[row]) {
						xmas = false
						break
					}

					if grid[rOffset][cOffset] != rune(word[chi]) {
						xmas = false
						break
					}

				}

				if xmas {
					sum++
				}
			}
		}
	}

	return sum
}

func part2(input []string) int {
	var sum int
	grid := make([][]rune, len(input))

	directions := [][2][2]int{
		{
			{-1, 1},
			{1, -1},
		},
		{
			{-1, -1},
			{1, 1},
		},
	}

	for k, line := range input {
		grid[k] = []rune(line)
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] != 'A' {
				continue
			}

			xmas := true

			for _, direction := range directions {
				r1Index := row + direction[0][0]
				c1Index := col + direction[0][1]

				r2Index := row + direction[1][0]
				c2Index := col + direction[1][1]

				if r1Index < 0 || r1Index >= len(grid) || c1Index < 0 || c1Index >= len(grid[row]) || r2Index < 0 || r2Index >= len(grid) || c2Index < 0 || c2Index >= len(grid[row]) {
					xmas = false
					break
				}

				if (grid[r1Index][c1Index] == 'M' && grid[r2Index][c2Index] == 'S') || (grid[r1Index][c1Index] == 'S' && grid[r2Index][c2Index] == 'M') {
					continue
				}

				xmas = false
				break
			}

			if xmas {
				sum++
			}
		}
	}

	return sum
}

func main() {
	nums := utils.GetInputs("input.txt")

	fmt.Print(
		"Solution AOC2024 Day4\n",
		"----------------------\n",
		fmt.Sprintf("Part 1: %v\n", part1(nums)),
		fmt.Sprintf("Part 2: %v\n", part2(nums)),
	)
}
