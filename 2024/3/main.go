package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/t0nyandre/adventofcode/utils"
)

func multiply(input [][]int) int {
	var sum int
	for _, v := range input {
		sum += v[0] * v[1]
	}
	return sum
}

func getNumbers(input [][]string) [][]int {
	numbers := make([][]int, 0)
	for _, v := range input {
		matchesSplit := strings.Split(v[0][4:len(v[0])-1], ",")
		x, _ := strconv.Atoi(matchesSplit[0])
		y, _ := strconv.Atoi(matchesSplit[1])
		numbers = append(numbers, []int{x, y})
	}
	return numbers
}

func part1(input []string) int {
	sum := 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	for _, v := range input {
		matches := re.FindAllStringSubmatch(v, -1)
		numbers := getNumbers(matches)
		sum += multiply(numbers)
	}
	return sum
}

func part2(input []string) int {
	var sum int
	enabled := true
	allEnabled := make([][]string, 0)
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	for _, v := range input {
		matches := re.FindAllStringSubmatch(v, -1)
		for _, i := range matches {
			if i[0] == "do()" {
				enabled = true
			} else if string(i[0]) == "don't()" {
				enabled = false
			} else if enabled {
				allEnabled = append(allEnabled, i)
			}
		}
	}
	numbers := getNumbers(allEnabled)

	sum += multiply(numbers)

	return sum
}

func main() {
	nums := utils.GetInputs("input.txt")

	fmt.Print(
		"Solution AOC2024 Day3\n",
		"----------------------\n",
		fmt.Sprintf("Part 1: %v\n", part1(nums)),
		fmt.Sprintf("Part 2: %v\n", part2(nums)),
	)
}
