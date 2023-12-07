package day07

import (
	"slices"
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

type Day07 struct{}

var FIVE_OF_KIND = 7
var FOUR_OF_KIND = 6
var FULL_HOUSE = 5
var THREE_OF_KIND = 4
var TWO_PAIR = 3
var ONE_PAIR = 2
var HIGH_CARD = 1

func (d Day07) Part1(input string) *util.Solution {
	var CARD_VALUES = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

	rawHands := util.Lines(input)
	hands := make([]Hand, len(rawHands))

	for i, raw := range rawHands {
		parts := strings.Split(raw, " ")

		hand := Hand{
			cards: parts[0],
			bid:   util.MustParseInt(parts[1]),
		}

		individualCards := 0
		cardCounts := map[string]int{}

		for i, card := range strings.Split(hand.cards, "") {
			count, hasCard := cardCounts[card]

			if !hasCard {
				cardCounts[card] = 1
				if !strings.ContainsAny(card, hand.cards[i+1:]) {
					individualCards++
				}

				continue
			}

			cardCounts[card] = count + 1
		}

		if len(cardCounts) == 1 {
			hand.kind = FIVE_OF_KIND
		} else if len(cardCounts) == 5 {
			hand.kind = HIGH_CARD
		} else if len(cardCounts) == 4 && individualCards == 3 {
			hand.kind = ONE_PAIR
		} else if len(cardCounts) == 3 && individualCards == 1 {
			hand.kind = TWO_PAIR
		} else if len(cardCounts) == 3 && individualCards == 2 {
			hand.kind = THREE_OF_KIND
		} else if len(cardCounts) == 2 && individualCards == 1 {
			hand.kind = FOUR_OF_KIND
		} else {
			hand.kind = FULL_HOUSE
		}

		hands[i] = hand
	}

	// sort hands asc according to rank
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
			valueA := slices.Index(CARD_VALUES, string(handA.cards[i]))
			valueB := slices.Index(CARD_VALUES, string(handB.cards[i]))
			if valueA == valueB {
				continue
			}

			return valueA < valueB
		}

		return true
	})

	totalWinnings := 0

	for rank, hand := range hands {
		totalWinnings += hand.bid * (rank + 1)
	}

	return util.NewSolution(1, totalWinnings)
}

func (d Day07) Part2(input string) *util.Solution {
	return util.NewSolution(2, -1)
}

type Hand struct {
	cards string
	kind  int
	bid   int
}
