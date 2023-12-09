package day09

import (
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day09 struct{}

func (d Day09) Part1(input string) *common.Solution {
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

	return common.NewSolution(1, nextValuesSum)
}

func (d Day09) Part2(input string) *common.Solution {
	sequences := util.Lines(input)

	previousValues := make([]int, len(sequences))

	for i, sequence := range sequences {
		rawValues := strings.Split(sequence, " ")
		values := make([]int, len(rawValues))

		for j, rawValue := range rawValues {
			values[j] = util.MustParseInt(rawValue)
		}

		previousValues[i] = findPreviousValue(values)
	}

	previousValuesSum := util.SumInts(previousValues)

	return common.NewSolution(2, previousValuesSum)
}

func findNextValue(values []int) int {
	equalsToZero := func(val int) bool { return val == 0 }
	if util.Every(values, equalsToZero) {
		return 0
	}

	gaps := make([]int, 0)

	for i := 0; i < len(values)-1; i++ {
		gaps = append(gaps, values[i+1]-values[i])
	}

	return util.LastElement(values) + findNextValue(gaps)
}

func findPreviousValue(values []int) int {
	equalsToZero := func(val int) bool { return val == 0 }
	if util.Every(values, equalsToZero) {
		return 0
	}

	gaps := make([]int, 0)

	for i := 0; i < len(values)-1; i++ {
		gaps = append(gaps, values[i+1]-values[i])
	}

	return values[0] - findPreviousValue(gaps)
}
