package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day04(input [][]string) (int, int) {
	var pairs int
	var overlap int

	for _, x := range input {
		elf1Start, elf1End := parseInt(x[0])
		elf2Start, elf2End := parseInt(x[1])
		if elf1Start >= elf2Start && elf1End <= elf2End {
			pairs++
		} else if elf2Start >= elf1Start && elf2End <= elf1End {
			pairs++
		}
		if elf1Start >= elf2Start && elf1Start <= elf2End {
			overlap++
		} else if elf2Start >= elf1Start && elf2Start <= elf1End {
			overlap++
		}
	}

	return pairs, overlap
}

func parseInt(input string) (int, int) {
	elf := strings.Split(input, "-")
	start, err := strconv.Atoi(elf[0])
	end, err := strconv.Atoi(elf[1])
	if err != nil {
		log.Fatalf("Could not convert string to int: %v", err)
	}
	return start, end
}

func main() {
	var input [][]string
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := scanner.Text()
		elves := strings.Split(pair, ",")
		input = append(input, elves)
	}

	fmt.Println(day04(input))
}
