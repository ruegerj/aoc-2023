package day$DAY_NR

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var dailyInput string

func TestMain(m *testing.M) {
	input, err := os.ReadFile(path.Join("..", "data", "$DAY_NR.txt"))

	if err != nil {
		fmt.Println("Failed to load input file", err)
		os.Exit(1)
	}

	dailyInput = string(input)

	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	expected := -1 // TODO: adapt
	solution := Day$DAY_NR{}.Part1(dailyInput)

	if solution.Result.(int) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}

func TestPart2(t *testing.T) {
	expected := -1 // TODO: adapt
	solution := Day$DAY_NR{}.Part2(dailyInput)

	if solution.Result.(int) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}