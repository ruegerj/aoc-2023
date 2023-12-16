package day15

import (
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
)

type Day15 struct{}

func (d Day15) Part1(input string) *common.Solution {
	totalHashSum := 0

	for _, part := range strings.Split(input, ",") {
		totalHashSum += hash(part)
	}

	return common.NewSolution(1, totalHashSum)
}

func (d Day15) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

func hash(value string) int {
	hash := 0

	for _, char := range value {
		hash += int(char)
		hash *= 17
		hash = hash % 256
	}

	return hash
}
