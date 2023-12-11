package util

func SumInts(items []int) int {
	sum := 0

	for _, i := range items {
		sum += i
	}

	return sum
}

func Chunks[T any](slice []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func LastElement[T comparable](slice []T) T {
	var lastElement T

	if slice == nil || len(slice) == 0 {
		return lastElement // return nil
	}

	return slice[len(slice)-1]
}

func RemoveIndex[T comparable](index int, slice []T) []T {
	if index == len(slice)-1 {
		return slice[:index]
	}

	return append(slice[:index], slice[index+1:]...)
}

func Every[T comparable](slice []T, predicate func(T) bool) bool {
	for _, value := range slice {
		if !predicate(value) {
			return false
		}
	}

	return true
}

func Flat[T any](slice [][]T) []T {
	flatSlice := make([]T, 0)

	for _, subSlice := range slice {
		flatSlice = append(flatSlice, subSlice...)
	}

	return flatSlice
}

func FirstOrDefault[T any](slice []T, predicate func(T) bool) T {
	var defaultValue T

	for _, value := range slice {
		if predicate(value) {
			return value
		}
	}

	return defaultValue
}

func Filter[T any](slice []T, filter func(T) bool) []T {
	filtered := make([]T, 0)

	for _, value := range slice {
		if filter(value) {
			filtered = append(filtered, value)
		}
	}

	return filtered
}

func Transpose[T any](matrix [][]T) [][]T {
	transposed := make([][]T, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			transposed[j] = append(transposed[j], matrix[i][j])
		}
	}
	return transposed
}
