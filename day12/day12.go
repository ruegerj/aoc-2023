package day12

import (
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const OPERATIONAL = "."
const DAMAGED = "#"
const UNKNOWN = "?"

type Day12 struct{}

func (d Day12) Part1(input string) *common.Solution {
	records := util.Lines(input)

	totalCombinations := 0

	for _, record := range records {
		parts := strings.Split(record, " ")
		pattern := parts[0]

		expectedDamaged := make([]int, 0)

		for _, expected := range strings.Split(parts[1], ",") {
			expectedDamaged = append(expectedDamaged, util.MustParseInt(expected))
		}

		foundCombinations := findPossibleCombinations(pattern, "", expectedDamaged)
		totalCombinations += foundCombinations
	}

	return common.NewSolution(1, totalCombinations)
}

func (d Day12) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

func findPossibleCombinations(pattern string, record string, expectedGroups []int) int {
	if len(pattern) == 0 {
		damagedGroups := make([]string, 0)
		springs := strings.Split(record, OPERATIONAL)

		for _, spring := range springs {
			if spring == "" {
				continue
			}

			damagedGroups = append(damagedGroups, spring)
		}

		if len(damagedGroups) != len(expectedGroups) {
			return 0
		}

		for i := 0; i < len(expectedGroups); i++ {
			if len(damagedGroups) <= i {
				return 0
			}

			actual := len(damagedGroups[i])
			expected := expectedGroups[i]

			if actual != expected {
				return 0
			}
		}

		return 1
	}

	current := pattern[0:1]
	unprocessed := ""

	if len(pattern) > 1 {
		unprocessed = pattern[1:]
	}

	if current == DAMAGED || current == OPERATIONAL {
		record += current
		return findPossibleCombinations(unprocessed, record, expectedGroups)
	}

	operationalVariation := findPossibleCombinations(unprocessed, record+OPERATIONAL, expectedGroups)
	damagedVariation := findPossibleCombinations(unprocessed, record+DAMAGED, expectedGroups)

	return operationalVariation + damagedVariation
}
