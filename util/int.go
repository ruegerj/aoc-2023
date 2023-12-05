package util

import (
	"strconv"
)

func MustParseInt(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	return number
}

func MustParseInt64(text string) int64 {
	num, err := strconv.ParseInt(text, 10, 64)

	if err != nil {
		panic(err)
	}

	return num
}

func Abs(number int) int {
	if number < 0 {
		number = number * -1
	}

	return number
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}
