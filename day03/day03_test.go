package day03

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var dailyInput string

func TestMain(m *testing.M) {
	input, err := os.ReadFile(path.Join("..", "data", "03.txt"))

	if err != nil {
		fmt.Println("Failed to load input file", err)
		os.Exit(1)
	}

	dailyInput = string(input)

	code := m.Run()
	os.Exit(code)
}

func TestPart1(t *testing.T) {
	expected := 554003
	solution := Day03{}.Part1(dailyInput)

	if solution.Result.(int) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}

func TestPart2(t *testing.T) {
	expected := 87263515
	solution := Day03{}.Part2(dailyInput)

	if solution.Result.(int) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}
