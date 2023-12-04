package day04

import (
	"regexp"
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

type Day04 struct{}

func (d Day04) Part1(input string) *util.Solution {
	partMatcher := regexp.MustCompile(`Card\s+\d+: (?P<winningNumbers>[\s|\d]+) \| (?P<ownNumbers>[\s|\d]+)`)

	totalPoints := 0

	for _, card := range util.Lines(input) {
		parts := util.MatchNamedSubgroups(partMatcher, card)

		winningNumbers := strings.Split(parts["winningNumbers"], " ")
		ownNumbers := strings.Split(parts["ownNumbers"], " ")

		ownWinningNumbers := make([]string, 0)

		for _, num := range ownNumbers {
			if !slices.Contains(winningNumbers, num) || strings.TrimSpace(num) == "" {
				continue
			}

			ownWinningNumbers = append(ownWinningNumbers, strings.TrimSpace(num))
		}

		points := 0

		for i := 0; i < len(ownWinningNumbers); i++ {
			if i == 0 {
				points = 1
				continue
			}

			points *= 2
		}

		totalPoints += points
	}

	return util.NewSolution(1, totalPoints)
}

func (d Day04) Part2(input string) *util.Solution {
	return util.NewSolution(2, -1)
}
