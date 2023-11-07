package day$DAY_NR

import "testing"

func TestPart1(t *testing.T) {
	expected := -1 // TODO: adapt
	solution := Day$DAY_NR{}.Part1()

	if solution.Result.(int) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}

func TestPart2(t *testing.T) {
	expected := -1 // TODO: adapt
	solution := Day$DAY_NR{}.Part2()

	if solution.Result.(int) != expected {
		t.Errorf("Expected %d, produced %d", expected, solution.Result)
	}
}