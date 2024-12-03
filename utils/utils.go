package utils

import (
	"bufio"
	"log"
	"math"
	"os"
)

func AlwaysPositive(left, right int) int {
	if (left - right) < 0 {
		return int(math.Abs(float64(left - right)))
	}
	return left - right
}

func GetInputs(path string) []string {
	var inputs []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	return inputs
}
