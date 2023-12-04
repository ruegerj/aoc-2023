package day03

import (
	"strconv"

	"github.com/ruegerj/aoc-2023/util"
)

type Day03 struct{}

func (d Day03) Part1(input string) *util.Solution {
	matrix := util.Matrix[string](input, "", func(s string) string { return s })

	partNumberSum := 0

	for i, row := range matrix {
		isPartNum := false
		num := ""

		for j, char := range row {
			if !isInt(char) && num != "" {
				if isPartNum {
					partNumberSum += util.MustParseInt(num)
				}

				isPartNum = false
				num = ""
				continue
			}

			if !isInt(char) {
				continue
			}

			num += char

			// right
			if len(matrix[i]) > j+1 && isSymbol(matrix[i][j+1]) {
				isPartNum = true
				continue
			}

			// left
			if j-1 >= 0 && isSymbol(matrix[i][j-1]) {
				isPartNum = true
				continue
			}

			// bottom
			if len(matrix) > i+1 && isSymbol(matrix[i+1][j]) {
				isPartNum = true
				continue
			}

			// top
			if i-1 >= 0 && isSymbol(matrix[i-1][j]) {
				isPartNum = true
				continue
			}

			// top right
			if i-1 >= 0 && len(matrix[i]) > j+1 && isSymbol(matrix[i-1][j+1]) {
				isPartNum = true
				continue
			}

			// top left
			if i-1 >= 0 && j-1 >= 0 && isSymbol(matrix[i-1][j-1]) {
				isPartNum = true
				continue
			}

			// bottom right
			if len(matrix) > i+1 && len(matrix[i]) > j+1 && isSymbol(matrix[i+1][j+1]) {
				isPartNum = true
				continue
			}

			// bottom lefts
			if len(matrix) > i+1 && j-1 >= 0 && isSymbol(matrix[i+1][j-1]) {
				isPartNum = true
				continue
			}
		}

		if isPartNum && num != "" {
			partNumberSum += util.MustParseInt(num)
		}
	}

	return util.NewSolution(1, partNumberSum)
}

func (d Day03) Part2(input string) *util.Solution {
	return util.NewSolution(2, -1)
}

func isInt(text string) bool {
	_, err := strconv.Atoi(text)
	return err == nil
}

func isSymbol(text string) bool {
	return !isInt(text) && text != "."
}
