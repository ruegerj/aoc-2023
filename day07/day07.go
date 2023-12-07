package day07

import (
	"slices"
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
	"golang.org/x/exp/maps"
)

type Day07 struct{}

const FIVE_OF_KIND = 7
const FOUR_OF_KIND = 6
const FULL_HOUSE = 5
const THREE_OF_KIND = 4
const TWO_PAIR = 3
const ONE_PAIR = 2
const HIGH_CARD = 1

const JOKER = "J"

func (d Day07) Part1(input string) *util.Solution {
	var CARD_VALUES = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	rawHands := util.Lines(input)
	hands := make([]Hand, len(rawHands))

	for i, raw := range rawHands {
		parts := strings.Split(raw, " ")

		cards := parts[0]
		bid := util.MustParseInt(parts[1])
		cardCounts, individualCards := analyzeCards(cards, false)

		hands[i] = Hand{
			cards: cards,
			bid:   bid,
			kind:  resolveHandKind(cardCounts, individualCards),
		}
	}

	sortHandsByRankAsc(hands, CARD_VALUES)
	totalWinnings := calcTotalWinnings(hands)

	return util.NewSolution(1, totalWinnings)
}

func (d Day07) Part2(input string) *util.Solution {
	var CARD_VALUES = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

	rawHands := util.Lines(input)
	hands := make([]Hand, len(rawHands))

	for i, raw := range rawHands {
		parts := strings.Split(raw, " ")

		cards := parts[0]
		bid := util.MustParseInt(parts[1])
		cardCounts, individualCards := analyzeCards(cards, true)
		jokerCount, hasJoker := cardCounts[JOKER]

		if hasJoker && len(cardCounts) > 1 {
			delete(cardCounts, JOKER)
			counts := maps.Values(cardCounts)
			sort.Slice(counts, func(i, j int) bool {
				return counts[i] > counts[j]
			})

			mostCommonCardCount := counts[0]
			mostCardsKey := util.KeyByValue(cardCounts, mostCommonCardCount)

			cardCounts[mostCardsKey] += jokerCount

			if mostCommonCardCount == 1 {
				individualCards--
			}
		}

		hands[i] = Hand{
			cards: cards,
			bid:   bid,
			kind:  resolveHandKind(cardCounts, individualCards),
		}
	}

	sortHandsByRankAsc(hands, CARD_VALUES)
	totalWinnings := calcTotalWinnings(hands)

	return util.NewSolution(2, totalWinnings)
}

func analyzeCards(cards string, ignoreJoker bool) (map[string]int, int) {
	cardCounts := map[string]int{}
	individualCards := 0

	for i, card := range strings.Split(cards, "") {
		count, hasCard := cardCounts[card]

		if !hasCard {
			cardCounts[card] = 1
			if strings.ContainsAny(card, cards[i+1:]) || (ignoreJoker && card == JOKER) {
				continue
			}

			individualCards++
		}

		cardCounts[card] = count + 1
	}

	return cardCounts, individualCards
}

func resolveHandKind(cardCounts map[string]int, individualCards int) int {
	if len(cardCounts) == 1 {
		return FIVE_OF_KIND
	}
	if len(cardCounts) == 5 {
		return HIGH_CARD
	}
	if len(cardCounts) == 4 && individualCards == 3 {
		return ONE_PAIR
	}
	if len(cardCounts) == 3 && individualCards == 1 {
		return TWO_PAIR
	}
	if len(cardCounts) == 3 && individualCards == 2 {
		return THREE_OF_KIND
	}
	if len(cardCounts) == 2 && individualCards == 1 {
		return FOUR_OF_KIND
	}

	return FULL_HOUSE
}

func sortHandsByRankAsc(hands []Hand, cardValues []string) {
	sort.Slice(hands, func(a, b int) bool {
		handA := hands[a]
		handB := hands[b]

		if handA.kind < handB.kind {
			return true
		}

		if handB.kind < handA.kind {
			return false
		}

		for i := 0; i < 5; i++ {
			valueA := slices.Index(cardValues, string(handA.cards[i]))
			valueB := slices.Index(cardValues, string(handB.cards[i]))
			if valueA == valueB {
				continue
			}

			return valueA < valueB
		}

		return true
	})
}

func calcTotalWinnings(hands []Hand) int {
	totalWinnings := 0

	for rank, hand := range hands {
		totalWinnings += hand.bid * (rank + 1)
	}

	return totalWinnings
}

type Hand struct {
	cards string
	kind  int
	bid   int
}
