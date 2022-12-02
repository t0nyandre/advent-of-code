package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Score struct {
	part1 int
	part2 int
}

func aocDay2(input *os.File) (int, int) {
	scorePart1 := 0
	scorePart2 := 0

	// Two players, multiple rounds, rock > scissors, scissors > paper, paper > rock
	// rock = 1, paper = 2, scissors = 3, loss = 0, draw = 3, win = 6
	points := map[string]Score{
		"A X": {4, 3},
		"B X": {1, 1},
		"C X": {7, 2},

		"A Y": {8, 4},
		"B Y": {5, 5},
		"C Y": {2, 6},

		"A Z": {3, 8},
		"B Z": {9, 9},
		"C Z": {6, 7},
	}

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		round := scanner.Text()
		scorePart1 += points[round].part1
		scorePart2 += points[round].part2
	}
	part1 := scorePart1
	part2 := scorePart2
	return part1, part2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(aocDay2(file))
}
