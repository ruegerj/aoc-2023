package day14

import (
	"slices"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const CUBE = "#"
const ROUND = "O"
const EMPTY = "."

type Day14 struct{}

func (d Day14) Part1(input string) *common.Solution {
	matrix := util.Matrix(input, "")
	transposedMatrix := util.Transpose(matrix)

	totalLoad := 0

	for _, row := range transposedMatrix {
		slices.Reverse(row)

		rollingRocks := make([]string, 0)

		for i, current := range row {
			if current == EMPTY {
				continue
			}

			if current == ROUND {
				rollingRocks = append(rollingRocks, current)
				continue
			}

			totalLoad += calcLoad(rollingRocks, i)
			rollingRocks = make([]string, 0)
		}

		totalLoad += calcLoad(rollingRocks, len(row))
	}

	return common.NewSolution(1, totalLoad)
}

func (d Day14) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

func calcLoad(rocks []string, maxLevel int) int {
	load := 0

	for i := 0; i < len(rocks); i++ {
		load += maxLevel
		maxLevel--
	}

	return load
}
