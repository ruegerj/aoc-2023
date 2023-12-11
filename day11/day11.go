package day11

import (
	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day11 struct{}

const SPACE = "."
const GALAXY = "#"
const RANGE_MULTIPLIER = 2

func (d Day11) Part1(input string) *common.Solution {
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

	totalRanges := 0
	seenPaths := make([]Path, 0)

	for _, galaxy := range galaxies {
		for i := 0; i < len(galaxies); i++ {
			next := galaxies[i]

			if galaxy == next {
				continue
			}

			pathCovered := util.Any(seenPaths, func(p Path) bool {
				return (p.from == galaxy && p.to == next) || (p.from == next && p.to == galaxy)
			})

			if pathCovered {
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

			deltaX := len(applicableCols)*RANGE_MULTIPLIER - len(applicableCols) + util.Abs(next.x-galaxy.x)
			deltaY := len(applicableRows)*RANGE_MULTIPLIER - len(applicableRows) + util.Abs(next.y-galaxy.y)

			seenPaths = append(seenPaths, Path{
				from: galaxy,
				to:   next,
			})
			effectiveRange := deltaY + deltaX
			totalRanges += effectiveRange
		}
	}

	return common.NewSolution(1, totalRanges)
}

func (d Day11) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

type Galaxy struct {
	x int
	y int
}

type Path struct {
	from *Galaxy
	to   *Galaxy
}
