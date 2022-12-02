package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func aocDay1(input *os.File) (int, int) {
	allElves := make([]int, 0)
	currentElf := make([]int, 0)

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			sum := 0
			for _, i := range currentElf {
				sum += i
			}
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

	sum := 0
	for _, i := range allElves[len(allElves)-3:] {
		sum += i
	}
	part2 := sum

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
