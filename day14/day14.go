package day14

import (
	"slices"
	"sort"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const CUBE = "#"
const ROUND = "O"
const EMPTY = "."

type Day14 struct{}

func (d Day14) Part1(input string) *common.Solution {
	matrix := util.Matrix(input, "")
	matrix = util.Transpose(matrix)

	for i, row := range matrix {
		matrix[i] = rollLeft(row)
	}

	northLoad := calcNorthLoad(matrix, true)

	return common.NewSolution(1, northLoad)
}

func (d Day14) Part2(input string) *common.Solution {
	matrix := util.Matrix(input, "")

	// the pattern shifts due to the rotations end up in a cycle
	// after 1000 rotations one can be certain that this cycle is reached due to the architecture of the input
	for i := 0; i < 1000; i++ {
		// north
		matrix = util.Transpose(matrix)
		for i, row := range matrix {
			matrix[i] = rollLeft(row)
		}

		// west
		matrix = util.Transpose(matrix)
		for i, row := range matrix {
			matrix[i] = rollLeft(row)
		}

		// south
		matrix = util.Transpose(matrix)
		for i, row := range matrix {
			matrix[i] = rollRight(row)
		}

		// east
		matrix = util.Transpose(matrix)
		for i, row := range matrix {
			matrix[i] = rollRight(row)
		}
	}

	northLoad := calcNorthLoad(matrix, false)

	return common.NewSolution(2, northLoad)
}

func rollRight(fields []string) []string {
	field := strings.Join(fields, "")
	parts := strings.Split(field, CUBE)

	for i, part := range parts {
		rocks := strings.Split(part, "")
		slices.Sort(rocks)
		parts[i] = strings.Join(rocks, "")
	}

	return strings.Split(strings.Join(parts, CUBE), "")
}

func rollLeft(fields []string) []string {
	field := strings.Join(fields, "")
	parts := strings.Split(field, CUBE)

	for i, part := range parts {
		rocks := strings.Split(part, "")
		sort.Sort(sort.Reverse(sort.StringSlice(rocks)))
		parts[i] = strings.Join(rocks, "")
	}

	return strings.Split(strings.Join(parts, CUBE), "")
}

func calcNorthLoad(matrix [][]string, transposed bool) int {
	northLoad := 0
	if !transposed {
		matrix = util.Transpose(matrix)
	}

	for _, row := range matrix {
		for i, rock := range row {
			if rock == ROUND {
				northLoad += len(row) - i
			}
		}
	}

	return northLoad
}
