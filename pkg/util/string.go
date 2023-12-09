package util

import "strings"

func Lines(text string) []string {
	return strings.Split(text, "\n")
}

func Matrix[T any](text string, colSeparator string, colTransformer func(string) T) [][]T {
	var matrix [][]T
	rows := Lines(text)

	for i := 0; i < len(rows); i++ {
		cols := strings.Split(rows[i], colSeparator)
		matrix = append(matrix, make([]T, len(cols)))

		for j := 0; j < len(cols); j++ {
			matrix[i][j] = colTransformer(cols[j])
		}
	}

	return matrix
}
