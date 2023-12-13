package day13

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day13 struct{}

func (d Day13) Part1(input string) *common.Solution {
	patterns := strings.Split(input, "\n\n")
	score := 0

	for i, pattern := range patterns {
		matrix := util.Matrix(pattern, "")
		horizontalMirror, hasHorizontal := findMirrorIndex(matrix)

		if hasHorizontal {
			score += horizontalMirror * 100
		}

		tMatrix := util.Transpose(matrix)
		verticalMirror, hasVertical := findMirrorIndex(tMatrix)

		if hasVertical {
			score += verticalMirror
		}

		if !hasHorizontal && !hasVertical {
			fmt.Println(i)
		}
	}

	return common.NewSolution(1, score)
}

func (d Day13) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

func findMirrorIndex(matrix [][]string) (int, bool) {
	candidates := make([]int, 0)

	for i, row := range matrix {
		if i == len(matrix)-1 {
			continue
		}

		if stringSlicesEqual(row, matrix[i+1]) {
			candidates = append(candidates, i)
		}
	}

	if len(candidates) == 0 {
		return 0, false
	}

	confirmedReflection := 0

	for _, candidate := range candidates {
		topMirror := candidate
		bottomMirror := candidate + 1

		positionsToCheck := topMirror
		deltaLowerBound := len(matrix) - 1 - bottomMirror

		if deltaLowerBound < topMirror {
			positionsToCheck = deltaLowerBound
		}

		lowerItems := matrix[bottomMirror+1 : bottomMirror+positionsToCheck+1]
		upperItems := util.DeepCopySlice(matrix[topMirror-positionsToCheck : topMirror])

		slices.Reverse(upperItems)
		successfullyVisited := 0

		for i := 0; i < positionsToCheck; i++ {
			lower := lowerItems[i]
			upper := upperItems[i]

			if !stringSlicesEqual(lower, upper) {
				break
			}

			successfullyVisited++
		}

		if successfullyVisited == positionsToCheck {
			confirmedReflection = candidate + 1
			break
		}
	}

	return confirmedReflection, confirmedReflection > 0
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	aStr := strings.Join(a, "")
	bStr := strings.Join(b, "")

	return aStr == bStr
}
