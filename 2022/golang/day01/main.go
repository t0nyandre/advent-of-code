package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/t0nyandre/advent-of-code/tree/main/2022/golang/utils"
)

func aocDay1(input *os.File) (int, int) {
	allElves := make([]int, 0)
	currentElf := make([]int, 0)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			sum := utils.Sum(currentElf)
			allElves = append(allElves, sum)
			currentElf = []int{}
		} else {
			cal, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			currentElf = append(currentElf, cal)
		}
	}

	sort.Ints(allElves)
	part1 := allElves[len(allElves)-1]
	part2 := utils.Sum(allElves[len(allElves)-3:])

	return part1, part2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	part1, part2 := aocDay1(file)
	fmt.Printf("Answer part 1: %v\nAnswer part 2: %v\n", part1, part2)
}
