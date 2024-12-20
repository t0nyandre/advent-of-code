package main

import (
	"os"
	"testing"

	"github.com/t0nyandre/adventofcode/utils"
)

var input []string
var exampleInput []string

func TestMain(m *testing.M) {
	code := 1
	defer func() { os.Exit(code) }()

	input = utils.GetInputs("input.txt")
	exampleInput = utils.GetInputs("example.txt")

	code = m.Run()
}

func Test_Day4(t *testing.T) {
	t.Run("Part1 Testcase", func(t *testing.T) {
		got := part1(exampleInput)
		// Please change the expected value
		expected := 18

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Part2 Testcase", func(t *testing.T) {
		got := part2(exampleInput)
		// Please change the expected value
		expected := 9

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Part1", func(t *testing.T) {
		got := part1(input)
		// Please change the expected value
		expected := 2718

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
	t.Run("Part2", func(t *testing.T) {
		got := part2(input)
		// Please change the expected value
		expected := 2046

		if got != expected {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
