package day06

import (
	"math"
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

	totalWinningVariations := calcTotalWinningVariationsOptimized(races)
	// totalWinningVariations := calcTotalWinningVariationsSemiBruteForce([]Race{race})

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

	totalWinningVariations := calcTotalWinningVariationsOptimized([]Race{race})
	// totalWinningVariations := calcTotalWinningVariationsSemiBruteForce([]Race{race})

	return util.NewSolution(2, totalWinningVariations)
}

func calcTotalWinningVariationsOptimized(races []Race) int {
	winningVariations := make([]int, len(races))

	// calculate the possible variations by solving the underlying quadratic equation
	for i, race := range races {
		a := float64(1)
		b := float64(race.time)
		c := float64(race.distance + 1)

		x1 := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
		x2 := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)

		winningVariations[i] = int(math.Floor(x1) - math.Ceil(x2) + 1)
	}

	totalWinningVariations := 1

	for _, variationCount := range winningVariations {
		totalWinningVariations *= variationCount
	}

	return totalWinningVariations
}

func calcTotalWinningVariationsSemiBruteForce(races []Race) int {
	winningVariations := make([]int, len(races))

	for i, race := range races {
		var lower, upper int

		for j := 1; j < race.time; j++ {
			effectiveDistance := j * (race.time - j)

			if effectiveDistance > race.distance {
				lower = j
				upper = race.time - j
				break
			}
		}

		winningVariations[i] = upper - lower + 1
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
