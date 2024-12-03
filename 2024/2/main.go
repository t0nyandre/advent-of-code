package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/t0nyandre/adventofcode/utils"
)

func isSafe(input []int) bool {
	// all increasing or all decreasing
	// differ by at least 1 or at most 3
	isIncreasing := true

	if input[0]-input[1] < 0 {
		isIncreasing = false
	}

	for i := 1; i < len(input); i++ {
		diff := input[i-1] - input[i]
		diffAbs := int(math.Abs(float64(diff)))
		if !(diffAbs >= 1 && diffAbs <= 3) {
			return false
		}
		if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			return false
		}
	}

	return true
}

func part1(input []string) int {
	var sum int
	for _, report := range input {
		reportSlice := strings.Fields(report)
		levels := make([]int, len(reportSlice))
		for k, v := range reportSlice {
			levels[k], _ = strconv.Atoi(v)
		}
		if isSafe(levels) {
			sum++
		}
	}

	return sum
}

func part2(input []string) int {
	var sum int
	for _, report := range input {
		reportSlice := strings.Fields(report)
		levels := make([]int, len(reportSlice))
		for k, v := range reportSlice {
			levels[k], _ = strconv.Atoi(v)
		}
		if isSafe(levels) {
			sum++
			continue
		}
		for i := 0; i < len(levels); i++ {
			modifiedLevels := []int{}
			modifiedLevels = append(modifiedLevels, levels[:i]...)
			modifiedLevels = append(modifiedLevels, levels[i+1:]...)
			if isSafe(modifiedLevels) {
				sum++
				break
			}
		}
	}

	return sum
}

func main() {
	reports := utils.GetInputs("input.txt")

	fmt.Println(part1(reports), part2(reports))
}
