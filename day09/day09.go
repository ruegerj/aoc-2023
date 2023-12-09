package day09

import (
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

type Day09 struct{}

func (d Day09) Part1(input string) *util.Solution {
	sequences := util.Lines(input)

	nextValues := make([]int, len(sequences))

	for i, sequence := range sequences {
		rawValues := strings.Split(sequence, " ")
		values := make([]int, len(rawValues))

		for j, rawValue := range rawValues {
			values[j] = util.MustParseInt(rawValue)
		}

		nextValues[i] = findNextValue(values)
	}

	nextValuesSum := util.SumInts(nextValues)

	return util.NewSolution(1, nextValuesSum)
}

func (d Day09) Part2(input string) *util.Solution {
	return util.NewSolution(2, -1)
}

func findNextValue(values []int) int {
	if util.Every(values, func(val int) bool { return val == 0 }) {
		return 0
	}

	gaps := make([]int, 0)

	for i := 0; i < len(values)-1; i++ {
		gaps = append(gaps, values[i+1]-values[i])
	}

	return util.LastElement(values) + findNextValue(gaps)
}
