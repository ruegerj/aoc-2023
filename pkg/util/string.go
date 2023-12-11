package util

import "strings"

func Lines(text string) []string {
	return strings.Split(text, "\n")
}

func Matrix(text string, colSeparator string) [][]string {
	var matrix [][]string
	rows := Lines(text)

	for i := 0; i < len(rows); i++ {
		cols := strings.Split(rows[i], colSeparator)
		matrix = append(matrix, make([]string, len(cols)))

		for j := 0; j < len(cols); j++ {
			matrix[i][j] = cols[j]
		}
	}

	return matrix
}
