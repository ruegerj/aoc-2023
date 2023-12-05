package day05

import (
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

type Day05 struct{}

func (d Day05) Part1(input string) *util.Solution {
	seeds, plan := parseSeedPlan(input)

	seedLocations := make([]int, len(seeds))

	for i, seed := range seeds {
		soil := resolve(seed, plan.seedToSoil)
		fertilizer := resolve(soil, plan.soilToFertilizer)
		water := resolve(fertilizer, plan.fertilizerToWater)
		light := resolve(water, plan.waterToLight)
		temperature := resolve(light, plan.lightToTemperature)
		humidity := resolve(temperature, plan.temperatureToHumidity)
		location := resolve(humidity, plan.humidityToLocation)

		seedLocations[i] = location
	}

	slices.Sort(seedLocations)

	return util.NewSolution(1, seedLocations[0])
}

func (d Day05) Part2(input string) *util.Solution {
	return util.NewSolution(2, -1)
}

func resolve(start int, ranges []Range) int {
	for _, r := range ranges {
		if start < r.source || start > r.source+r.delta {
			continue
		}

		offset := start - r.source
		return r.destination + offset
	}

	return start
}

func parseSeedPlan(input string) ([]int, *SeedPlan) {
	parts := strings.Split(input, "\n\n")
	plan := &SeedPlan{}

	// parse seeds
	seeds := make([]int, 0)
	for _, seed := range strings.Split(strings.Split(parts[0], ": ")[1], " ") {
		seeds = append(seeds, util.MustParseInt(seed))
	}

	// parse all mapping tables
	seedToSoil := util.Lines(parts[1])[1:]
	plan.seedToSoil = parseRanges(seedToSoil)

	soilToFertilizer := util.Lines(parts[2])[1:]
	plan.soilToFertilizer = parseRanges(soilToFertilizer)

	fertilizerToWater := util.Lines(parts[3])[1:]
	plan.fertilizerToWater = parseRanges(fertilizerToWater)

	waterToLight := util.Lines(parts[4])[1:]
	plan.waterToLight = parseRanges(waterToLight)

	lightToTemperature := util.Lines(parts[5])[1:]
	plan.lightToTemperature = parseRanges(lightToTemperature)

	temperatureToHumidity := util.Lines(parts[6])[1:]
	plan.temperatureToHumidity = parseRanges(temperatureToHumidity)

	humidityToLocation := util.Lines(parts[7])[1:]
	plan.humidityToLocation = parseRanges(humidityToLocation)

	return seeds, plan
}

func parseRanges(rawRanges []string) []Range {
	ranges := make([]Range, len(rawRanges))

	for i, rawRange := range rawRanges {
		parts := strings.Split(rawRange, " ")

		ranges[i] = Range{
			source:      util.MustParseInt(parts[1]),
			destination: util.MustParseInt(parts[0]),
			delta:       util.MustParseInt(parts[2]),
		}
	}

	return ranges
}

type SeedPlan struct {
	seedToSoil            []Range
	soilToFertilizer      []Range
	fertilizerToWater     []Range
	waterToLight          []Range
	lightToTemperature    []Range
	temperatureToHumidity []Range
	humidityToLocation    []Range
}

type Range struct {
	source      int
	destination int
	delta       int
}
