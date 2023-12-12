package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const OPERATIONAL = '.'
const DAMAGED = '#'
const UNKNOWN = '?'
const NO_GROUP = -1

type Day12 struct{}

func (d Day12) Part1(input string) *common.Solution {
	records := util.Lines(input)

	totalCombinations := 0
	memo := make(map[string]int, 0)

	for _, record := range records {
		parts := strings.Split(record, " ")
		pattern := parts[0]
		groups := parts[1]
		damagedSprings := util.ToIntSlice(strings.Split(groups, ","))

		foundCombinations := findCombinationsMemo(pattern, damagedSprings, NO_GROUP, memo)
		totalCombinations += foundCombinations
	}

	return common.NewSolution(1, totalCombinations)
}

func (d Day12) Part2(input string) *common.Solution {
	records := util.Lines(input)

	totalCombinations := 0
	memo := make(map[string]int, 0)

	for _, record := range records {
		parts := strings.Split(record, " ")
		pattern, groups := unfold(parts[0], parts[1])
		damagedSprings := util.ToIntSlice(strings.Split(groups, ","))

		foundCombinations := findCombinationsMemo(pattern, damagedSprings, NO_GROUP, memo)
		totalCombinations += foundCombinations
	}

	return common.NewSolution(2, totalCombinations)
}

func findCombinationsMemo(pattern string, groups []int, currentGroup int, memo map[string]int) int {
	memoedCombinations, foundEntry := memo[createKey(pattern, groups, currentGroup)]

	if foundEntry {
		return memoedCombinations
	}

	// all chars processed, group conditions met -> match
	if len(pattern) == 0 && len(groups) == 0 && currentGroup <= 0 {
		return 1
	}

	// all chars processed, group conditions not met -> no match
	if len(pattern) == 0 {
		return 0
	}

	combinations := 0

	switch pattern[0] {
	case OPERATIONAL:
		// current group not processed till end but interrupted with operational spring -> no match
		if currentGroup > 0 {
			return 0
		}

		combinations = findCombinationsMemo(pattern[1:], groups, NO_GROUP, memo)
	case DAMAGED:
		// current group or all groups are already processed -> no match
		if currentGroup == 0 || (len(groups) == 0 && currentGroup == NO_GROUP) {
			return 0
		}

		// pop next group to check
		if currentGroup == NO_GROUP {
			currentGroup = groups[0]
			groups = groups[1:]
		}

		combinations = findCombinationsMemo(pattern[1:], groups, currentGroup-1, memo)
	case UNKNOWN:
		operationalCombinations := findCombinationsMemo(string(OPERATIONAL)+pattern[1:], groups, currentGroup, memo)
		damagedCombinations := findCombinationsMemo(string(DAMAGED)+pattern[1:], groups, currentGroup, memo)
		combinations = operationalCombinations + damagedCombinations
	}

	memo[createKey(pattern, groups, currentGroup)] = combinations
	return combinations
}

func createKey(pattern string, remainingGroups []int, currentGroup int) string {
	convertedGroups := make([]string, len(remainingGroups))

	for i, group := range remainingGroups {
		convertedGroups[i] = strconv.Itoa(group)
	}

	return fmt.Sprintf("%s:%s:%d", pattern, strings.Join(convertedGroups, ","), currentGroup)
}

func unfold(pattern string, groups string) (string, string) {
	unfoldedPattern := pattern
	unfoldedGroups := groups

	for i := 1; i < 5; i++ {
		unfoldedPattern += string(UNKNOWN) + pattern
		unfoldedGroups += "," + groups
	}

	return unfoldedPattern, unfoldedGroups
}
