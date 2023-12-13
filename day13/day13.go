package day13

import (
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day13 struct{}

func (d Day13) Part1(input string) *common.Solution {
	patterns := strings.Split(input, "\n\n")
	score := 0

	for _, pattern := range patterns {
		matrix := util.Matrix(pattern, "")
		horizontalMirror, hasHorizontal := findMirrorIndex(matrix, 0)

		if hasHorizontal {
			score += horizontalMirror * 100
		}

		tMatrix := util.Transpose(matrix)
		verticalMirror, hasVertical := findMirrorIndex(tMatrix, 0)

		if hasVertical {
			score += verticalMirror
		}
	}

	return common.NewSolution(1, score)
}

func (d Day13) Part2(input string) *common.Solution {
	patterns := strings.Split(input, "\n\n")
	score := 0

	for _, pattern := range patterns {
		matrix := util.Matrix(pattern, "")
		horizontalMirror, hasHorizontal := findMirrorIndex(matrix, 1)

		if hasHorizontal {
			score += horizontalMirror * 100
		}

		tMatrix := util.Transpose(matrix)
		verticalMirror, hasVertical := findMirrorIndex(tMatrix, 1)

		if hasVertical {
			score += verticalMirror
		}
	}

	return common.NewSolution(2, score)
}

func findMirrorIndex(matrix [][]string, maxSmudges int) (int, bool) {
	candidates := make([]int, 0)

	for i, row := range matrix {
		if i == len(matrix)-1 {
			continue
		}

		current := row
		next := matrix[i+1]

		currentConcat := strings.Join(current, "")
		nextConcat := strings.Join(next, "")

		smudgeCount := countSmudges(currentConcat, nextConcat)

		if smudgeCount <= maxSmudges {
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

		lowerItems := matrix[bottomMirror : bottomMirror+positionsToCheck+1]
		upperItems := util.DeepCopySlice(matrix[topMirror-positionsToCheck : topMirror+1])
		positionsToCheck++

		slices.Reverse(upperItems)

		successfullyVisited := 0
		smudgesCorrected := 0

		for i := 0; i < positionsToCheck; i++ {
			lower := strings.Join(lowerItems[i], "")
			upper := strings.Join(upperItems[i], "")

			smudgeCount := countSmudges(lower, upper)

			if smudgeCount == 1 && smudgesCorrected < maxSmudges {
				smudgesCorrected++
			}

			if smudgeCount <= maxSmudges {
				successfullyVisited++
			}
		}

		if successfullyVisited == positionsToCheck && smudgesCorrected == maxSmudges {
			confirmedReflection = candidate + 1
			break
		}
	}

	return confirmedReflection, confirmedReflection > 0
}

func countSmudges(a, b string) int {
	smudges := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			smudges++
		}
	}

	return smudges
}
