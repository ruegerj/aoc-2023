package day04

import (
	"regexp"
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day04 struct{}

var CARD_MATCHER = regexp.MustCompile(`Card\s+\d+: (?P<winningNumbers>[\s|\d]+) \| (?P<ownNumbers>[\s|\d]+)`)

func (d Day04) Part1(input string) *common.Solution {
	totalPoints := 0

	cards := util.Lines(input)

	for i := 0; i < len(cards); i++ {
		winningNumbers := resolveWinningNumbers(i, cards)

		points := 0

		for i := 0; i < len(winningNumbers); i++ {
			if i == 0 {
				points = 1
				continue
			}

			points *= 2
		}

		totalPoints += points
	}

	return common.NewSolution(1, totalPoints)
}

func (d Day04) Part2(input string) *common.Solution {
	totalCards := 0

	cards := util.Lines(input)
	cardCount := map[int]int{}
	wonCopies := map[int]int{}

	for i := 0; i < len(cards); i++ {
		wonCopies[i] = -1
	}

	for i := 0; i < len(cards); i++ {
		cardCount[i] += 1
		resolveWonCards(i, cards, cardCount, wonCopies)
		totalCards += cardCount[i]
	}

	return common.NewSolution(2, totalCards)
}

func resolveWonCards(cardIndex int, cards []string, cardCount map[int]int, wonCopies map[int]int) {
	if cardIndex < 0 || cardIndex >= len(cards) {
		return
	}

	if wonCopies[cardIndex] == -1 {
		winningNumbers := resolveWinningNumbers(cardIndex, cards)
		wonCopies[cardIndex] = len(winningNumbers)
	}

	for i := 1; i <= wonCopies[cardIndex]; i++ {
		cardCount[cardIndex+i] += 1
		resolveWonCards(cardIndex+i, cards, cardCount, wonCopies)
	}
}

func resolveWinningNumbers(cardIndex int, cards []string) []string {
	parts := util.MatchNamedSubgroups(CARD_MATCHER, cards[cardIndex])

	winningNumbers := strings.Split(parts["winningNumbers"], " ")
	ownNumbers := strings.Split(parts["ownNumbers"], " ")
	ownWinningNumbers := make([]string, 0)

	for _, original := range ownNumbers {
		if !slices.Contains(winningNumbers, original) || strings.TrimSpace(original) == "" {
			continue
		}

		ownWinningNumbers = append(ownWinningNumbers, strings.TrimSpace(original))
	}

	return ownWinningNumbers
}
