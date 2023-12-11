package day11

import (
	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
	"golang.org/x/exp/maps"
)

type Day11 struct{}

const SPACE = "."
const GALAXY = "#"

func (d Day11) Part1(input string) *common.Solution {
	const RANGE_MULTIPLIER = 2
	galaxies, rowsWithModifier, colsWithModifier := parseUniverse(input)
	totalRanges := calcShortestPathSum(galaxies, rowsWithModifier, colsWithModifier, RANGE_MULTIPLIER)

	return common.NewSolution(1, totalRanges)
}

func (d Day11) Part2(input string) *common.Solution {
	const RANGE_MULTIPLIER = 1000000
	galaxies, rowsWithModifier, colsWithModifier := parseUniverse(input)
	totalRanges := calcShortestPathSum(galaxies, rowsWithModifier, colsWithModifier, RANGE_MULTIPLIER)

	return common.NewSolution(2, totalRanges)
}

func calcShortestPathSum(galaxies []*Galaxy, rowsWithModifier []int, colsWithModifier []int, rangeMultiplier int) int {
	pathLengths := map[Path]int{}

	for _, galaxy := range galaxies {
		for i := 0; i < len(galaxies); i++ {
			next := galaxies[i]

			if galaxy == next {
				continue
			}

			pathA := Path{from: galaxy, to: next}
			pathB := Path{from: next, to: galaxy}

			_, hasVarA := pathLengths[pathA]
			_, hasVarB := pathLengths[pathB]

			if hasVarA || hasVarB {
				continue
			}

			applicableRows := make([]int, 0)
			for _, row := range rowsWithModifier {
				if (galaxy.y < row && row < next.y) || (galaxy.y > row && row > next.y) {
					applicableRows = append(applicableRows, row)
				}
			}

			applicableCols := make([]int, 0)
			for _, col := range colsWithModifier {
				if (galaxy.x < col && col < next.x) || (galaxy.x > col && col > next.x) {
					applicableCols = append(applicableCols, col)
				}
			}

			deltaX := len(applicableCols)*rangeMultiplier - len(applicableCols) + util.Abs(next.x-galaxy.x)
			deltaY := len(applicableRows)*rangeMultiplier - len(applicableRows) + util.Abs(next.y-galaxy.y)

			effectiveRange := deltaY + deltaX

			pathLengths[pathA] = effectiveRange
			pathLengths[pathB] = 0
		}
	}

	return util.SumInts(maps.Values(pathLengths))
}

func parseUniverse(input string) ([]*Galaxy, []int, []int) {
	universe := util.Matrix(input, "")
	rowsWithModifier := make([]int, 0)

	for i, row := range universe {
		if util.Every(row, func(value string) bool { return value == SPACE }) {
			rowsWithModifier = append(rowsWithModifier, i)
		}
	}

	tUniverse := util.Transpose(universe)
	colsWithModifier := make([]int, 0)

	for i, col := range tUniverse {
		if util.Every(col, func(value string) bool { return value == SPACE }) {
			colsWithModifier = append(colsWithModifier, i)
		}
	}

	galaxies := make([]*Galaxy, 0)

	for y, row := range universe {
		for x, value := range row {
			if value == GALAXY {
				galaxies = append(galaxies, &Galaxy{
					x: x,
					y: y,
				})
			}
		}
	}

	return galaxies, rowsWithModifier, colsWithModifier
}

type Galaxy struct {
	x int
	y int
}

type Path struct {
	from *Galaxy
	to   *Galaxy
}
