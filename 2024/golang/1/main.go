package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/t0nyandre/adventofcode/utils"
)

func part1(left, right []int) int {
	var sum int
	for k, i := range left {
		sum = sum + utils.AlwaysPositive(i, right[k])
	}

	return sum
}

func part2(left, right []int) int {
	var sum int
	countMap := make(map[int]int)
	for _, i := range right {
		countMap[i]++
	}
	for i := 1; i < len(left); i++ {
		sum = sum + (left[i] * countMap[left[i]])
	}
	return sum
}

func main() {
	var left []int
	var right []int
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		IDs := strings.Fields(scanner.Text())
		leftInt, err := strconv.Atoi(IDs[0])
		if err != nil {
			log.Fatal(err)
		}
		rightInt, err := strconv.Atoi(IDs[1])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	sort.Ints(left)
	sort.Ints(right)

	fmt.Println(part1(left, right), part2(left, right))
}
