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

func Abs64(number int64) int64 {
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

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func LCM(a, b int, ints ...int) int {
	result := a * b / GCD(a, b)

	for _, num := range ints {
		result = LCM(result, num)
	}

	return result
}
