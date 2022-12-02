package main

import (
	"log"
	"os"
)

func aocDay2(input *os.File) (int, int) {
	part1 := 0
	part2 := 0
	return part1, part2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
