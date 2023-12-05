package day05

import (
	"slices"
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

type Day05 struct{}

func (d Day05) Part1(input string) *util.Solution {
	seeds, rangeMaps := parseSeedPlan(input)

	seedLocations := make([]int64, len(seeds))

	for i, seed := range seeds {
		location := seed

		for _, rangeMap := range rangeMaps {
			for _, mapping := range rangeMap {
				// check if range is suitable for mapping
				if location >= mapping.from && location <= mapping.to {
					location += mapping.delta
					break
				}
			}

			// value remains the same -> no suitable range found
		}

		seedLocations[i] = location
	}

	slices.Sort(seedLocations)

	return util.NewSolution(1, seedLocations[0])
}

func (d Day05) Part2(input string) *util.Solution {
	seedPairs, rangeMaps := parseSeedPlan(input)

	// generate seed ranges from seed pairs
	locationRanges := make([]Range, len(seedPairs)/2)

	for i := 0; i < len(seedPairs); i += 2 {
		startSeed := seedPairs[i]
		seedCount := seedPairs[i+1]

		locationRanges[i/2] = Range{
			from: startSeed,
			to:   startSeed + seedCount - 1,
		}
	}

	for _, rangeMap := range rangeMaps {
		sort.SliceStable(rangeMap, func(i, j int) bool {
			return rangeMap[i].from < rangeMap[j].from
		})

		mappedRanges := make([]Range, 0)

		for _, locationRange := range locationRanges {
			tmpRange := locationRange

			// iterate over all asc sorted mappings of the current stage (range map)
			for _, mapping := range rangeMap {
				// range starts left of the mapping bounds -> fit the gap around the mapping [r1, mapping, (r2)]
				if tmpRange.from < mapping.from {
					newRange := Range{
						from: tmpRange.from,
						to:   util.MinInt64(tmpRange.to, mapping.from-1),
					}
					mappedRanges = append(mappedRanges, newRange)
					tmpRange.from = mapping.from

					// if range is fully covered -> continue with next mapping
					if tmpRange.from >= tmpRange.to {
						break
					}
				}

				// range is inside of the mapping bounds -> transform it according to the mapping's delta
				if tmpRange.from <= mapping.to {
					newRange := Range{
						from: tmpRange.from + mapping.delta,
						to:   util.MinInt64(tmpRange.to, mapping.to) + mapping.delta,
					}
					mappedRanges = append(mappedRanges, newRange)
					tmpRange.from = mapping.to + 1

					// if range is fully covered -> continue with next mapping
					if tmpRange.from >= tmpRange.to {
						break
					}
				}
			}

			// no mapping covered the range -> re-use range 1:1
			if tmpRange.from < tmpRange.to {
				mappedRanges = append(mappedRanges, tmpRange)
			}
		}

		// update ranges with current stage
		locationRanges = mappedRanges
	}

	// the current locationRanges hold all ranges in which a location from a seed in the initial seed ranges could lay
	// -> take range with lowest from = nearest possible location
	sort.Slice(locationRanges, func(i, j int) bool {
		return locationRanges[i].from < locationRanges[j].from
	})

	return util.NewSolution(2, locationRanges[0].from)
}

func parseSeedPlan(input string) ([]int64, [][]RangeMap) {
	parts := strings.Split(input, "\n\n")

	// parse seeds
	seeds := make([]int64, 0)
	for _, seed := range strings.Split(strings.Split(parts[0], ": ")[1], " ") {
		seeds = append(seeds, util.MustParseInt64(seed))
	}

	// parse maps
	rangeMaps := make([][]RangeMap, len(parts)-1)
	for i, part := range parts[1:] {
		rawRanges := util.Lines(part)[1:]
		ranges := make([]RangeMap, len(rawRanges))

		for j, rawRange := range rawRanges {
			parts := strings.Split(rawRange, " ")

			dest := util.MustParseInt64(parts[0])
			source := util.MustParseInt64(parts[1])
			length := util.MustParseInt64(parts[2])

			ranges[j] = RangeMap{
				from:  source,
				to:    source + length - 1,
				delta: dest - source}
		}

		rangeMaps[i] = ranges
	}

	return seeds, rangeMaps
}

type RangeMap struct {
	from int64
	to   int64
	// value which should be used to transform a value inside the range bounds
	delta int64
}

type Range struct {
	from int64
	to   int64
}
