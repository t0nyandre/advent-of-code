package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(input []string) int {
	var sum int

	return sum
}

func part2(input []string) int {
	var sum int

	return sum
}

func main() {
	var input []string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(part1(input), part2(input))
}
