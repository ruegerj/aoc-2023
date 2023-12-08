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
