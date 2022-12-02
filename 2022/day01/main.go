package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func getReeindeers(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	reindeers := make([][]int, 0)
	curr := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		t, err := strconv.Atoi(line)
		if err != nil {
			reindeers = append(reindeers, curr)
			curr = make([]int, 0)
		} else {
			curr = append(curr, t)
		}
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	reindeers = append(reindeers, curr)
	return reindeers
}

func sumCalories(reindeers [][]int) []int {
	cal := make([]int, 0)
	for _, reindeer := range reindeers {
		sum := 0
		for _, i := range reindeer {
			sum += i
		}
		cal = append(cal, sum)
	}
	sort.Slice(cal, func(i, j int) bool {
		return cal[i] > cal[j]
	})

	return cal
}

func main() {
	reindeers := getReeindeers("input.txt")
	cals := sumCalories(reindeers)
	fmt.Println(cals[0])
	total := 0
	for _, i := range cals[:3] {
		total += i
	}
	fmt.Println(total)
}
