package day03

import (
	"strconv"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day03 struct{}

func (d Day03) Part1(input string) *common.Solution {
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

	return common.NewSolution(1, partNumberSum)
}

func (d Day03) Part2(input string) *common.Solution {
	matrix := util.Matrix[string](input, "", func(s string) string { return s })

	cogGearsLookup := map[Cog][]int{}

	for i, row := range matrix {
		var cog Cog
		isGear := false
		num := ""

		for j, char := range row {
			if !isInt(char) && num != "" {
				if isGear {
					cogGearsLookup[cog] = append(cogGearsLookup[cog], util.MustParseInt(num))
				}

				isGear = false
				cog = Cog{}
				num = ""
				continue
			}

			if !isInt(char) {
				continue
			}

			num += char

			// right
			if len(matrix[i]) > j+1 && isCogSymbol(matrix[i][j+1]) {
				isGear = true
				cog = Cog{x: j + 1, y: i}
				continue
			}

			// left
			if j-1 >= 0 && isCogSymbol(matrix[i][j-1]) {
				isGear = true
				cog = Cog{x: j - 1, y: i}
				continue
			}

			// bottom
			if len(matrix) > i+1 && isCogSymbol(matrix[i+1][j]) {
				isGear = true
				cog = Cog{x: j, y: i + 1}
				continue
			}

			// top
			if i-1 >= 0 && isCogSymbol(matrix[i-1][j]) {
				isGear = true
				cog = Cog{x: j, y: i - 1}
				continue
			}

			// top right
			if i-1 >= 0 && len(matrix[i]) > j+1 && isCogSymbol(matrix[i-1][j+1]) {
				isGear = true
				cog = Cog{x: j + 1, y: i - 1}
				continue
			}

			// top left
			if i-1 >= 0 && j-1 >= 0 && isCogSymbol(matrix[i-1][j-1]) {
				isGear = true
				cog = Cog{x: j - 1, y: i - 1}
				continue
			}

			// bottom right
			if len(matrix) > i+1 && len(matrix[i]) > j+1 && isCogSymbol(matrix[i+1][j+1]) {
				isGear = true
				cog = Cog{x: j + 1, y: i + 1}
				continue
			}

			// bottom lefts
			if len(matrix) > i+1 && j-1 >= 0 && isCogSymbol(matrix[i+1][j-1]) {
				isGear = true
				cog = Cog{x: j - 1, y: i + 1}
				continue
			}
		}

		if isGear && num != "" {
			cogGearsLookup[cog] = append(cogGearsLookup[cog], util.MustParseInt(num))
		}
	}

	gearRatioSum := 0

	for _, gears := range cogGearsLookup {
		if len(gears) != 2 {
			continue
		}

		gearRatioSum += gears[0] * gears[1]
	}

	return common.NewSolution(2, gearRatioSum)
}

func isInt(text string) bool {
	_, err := strconv.Atoi(text)
	return err == nil
}

func isSymbol(text string) bool {
	return !isInt(text) && text != "."
}

func isCogSymbol(text string) bool {
	return !isInt(text) && text == "*"
}

type Cog struct {
	x int
	y int
}
