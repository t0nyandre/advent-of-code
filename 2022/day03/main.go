package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/t0nyandre/advent-of-code/utils"
)

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func aocDay3(input *os.File) (int, int) {
	sum := 0
	compartment1 := []string{}
	compartment2 := []string{}
	pri := []string{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		rucksack := scanner.Text()
		length := len(rucksack)
		compartment1 = strings.Split(rucksack[:(length/2)], "")
		compartment2 = strings.Split(rucksack[(length/2):], "")
		rucksackPri := []string{}
		for _, x := range compartment1 {
			if utils.Contains(compartment2, x) {
				if len(rucksackPri) < 1 || rucksackPri[len(rucksackPri)-1] != x {
					rucksackPri = append(rucksackPri, x)
				}
			}
		}
		pri = append(pri, rucksackPri...)
	}

	for _, i := range pri {
		sum += strings.Index(priority, i) + 1
	}

	part1 := sum
	part2 := 0
	return part1, part2
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(aocDay3(file))
}
