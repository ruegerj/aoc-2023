package util

import (
	"fmt"
	"os"
	"strconv"
)

func MustParseInt(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Failed to parse input to integer")
		os.Exit(1)
	}

	return number
}

func Abs(number int) int {
	if number < 0 {
		number = number * -1
	}

	return number
}
