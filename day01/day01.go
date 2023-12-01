package day01

import (
	"strconv"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

var TEXT_TO_DIGIT = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type Day01 struct{}

func (d Day01) Part1(input string) *util.Solution {
	total := 0

	for _, value := range util.Lines(input) {
		digits := make([]string, 0)

		for _, char := range strings.Split(value, "") {
			_, err := strconv.Atoi(char)

			if err == nil {
				digits = append(digits, char)
			}

		}

		score := digits[0] + util.LastElement(digits)
		total += util.MustParseInt(score)
	}

	return util.NewSolution(1, total)
}

func (d Day01) Part2(input string) *util.Solution {
	total := 0

	for _, value := range util.Lines(input) {
		digits := make([]string, 0)
		chars := strings.Split(value, "")

		for i, char := range chars {
			digit, found := tryLookaheadDigitParse(i, chars)

			if found {
				digits = append(digits, digit)
				continue
			}

			_, err := strconv.Atoi(char)

			if err == nil {
				digits = append(digits, char)
			}
		}

		score := digits[0] + util.LastElement(digits)
		total += util.MustParseInt(score)
	}

	return util.NewSolution(2, total)
}

func tryLookaheadDigitParse(pos int, chars []string) (string, bool) {
	textLength := len(chars)
	wordCandidates := []string{}

	if pos+4 < textLength {
		wordCandidates = append(wordCandidates, strings.Join(chars[pos:pos+5], ""))
	}

	if pos+3 < textLength {
		wordCandidates = append(wordCandidates, strings.Join(chars[pos:pos+4], ""))
	}

	if pos+2 < textLength {
		wordCandidates = append(wordCandidates, strings.Join(chars[pos:pos+3], ""))
	}

	for _, candidate := range wordCandidates {
		digit, exits := TEXT_TO_DIGIT[candidate]

		if exits {
			return digit, true
		}
	}

	return "", false
}
