package day06

import (
	"regexp"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

type Day06 struct{}

var NUMBER_MATCHER = regexp.MustCompile(`\d+`)

func (d Day06) Part1(input string) *util.Solution {
	parts := util.Lines(input)
	times := NUMBER_MATCHER.FindAllString(parts[0], -1)
	distances := NUMBER_MATCHER.FindAllString(parts[1], -1)

	races := make([]Race, len(times))

	for i, time := range times {
		races[i] = Race{
			time:     util.MustParseInt(time),
			distance: util.MustParseInt(distances[i]),
		}
	}

	totalWinningVariations := calcTotalWinningVariations(races)

	return util.NewSolution(1, totalWinningVariations)
}

func (d Day06) Part2(input string) *util.Solution {
	parts := util.Lines(input)
	times := NUMBER_MATCHER.FindAllString(parts[0], -1)
	distances := NUMBER_MATCHER.FindAllString(parts[1], -1)

	race := Race{
		time:     util.MustParseInt(strings.Join(times, "")),
		distance: util.MustParseInt(strings.Join(distances, "")),
	}

	totalWinningVariations := calcTotalWinningVariations([]Race{race})

	return util.NewSolution(2, totalWinningVariations)
}

func calcTotalWinningVariations(races []Race) int {
	winningVariations := make([]int, 0)

	for _, race := range races {
		variations := 0

		// time is exclusive top & bottom
		for i := 1; i < race.time; i++ {
			effectiveDistance := i * (race.time - i)

			if effectiveDistance > race.distance {
				variations++
			}
		}

		winningVariations = append(winningVariations, variations)
	}

	totalWinningVariations := 1

	for _, variationCount := range winningVariations {
		totalWinningVariations *= variationCount
	}

	return totalWinningVariations
}

type Race struct {
	time     int
	distance int
}
