package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/t0nyandre/advent-of-code/tree/main/2022/golang/utils"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func part1(input []string) int {
	var (
		sum          int = 0
		compartment1 []string
		compartment2 []string
	)

	for _, rucksack := range input {
		items := []string{}
		compartment1 = strings.Split(rucksack[:(len(rucksack)/2)], "")
		compartment2 = strings.Split(rucksack[(len(rucksack)/2):], "")

		for _, x := range compartment1 {
			if utils.Contains(compartment2, x) {
				if len(items) < 1 || items[len(items)-1] != x {
					items = append(items, x)
				}
			}
		}
		sum += strings.Index(priority, items[len(items)-1]) + 1
	}

	return sum
}

func part2(input []string) int {
	sum := 0

	elves := utils.Chunks(input, 3)
	b := []string{}
	for _, i := range elves {
		b = strings.Split(i[0], "")
		for _, j := range b {
			if utils.Contains(strings.Split(i[1], ""), j) && utils.Contains(strings.Split(i[2], ""), j) {
				sum += strings.Index(priority, j) + 1
				break
			}
		}
	}

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
