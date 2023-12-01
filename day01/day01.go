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
	values := util.Lines(input)

	total := 0

	for _, value := range values {
		d1 := ""
		d2 := ""

		for _, c := range strings.Split(value, "") {
			_, err := strconv.Atoi(c)

			if err == nil && d1 == "" {
				d1 = c
				d2 = c
				continue
			}

			if err == nil {
				d2 = c
			}
		}

		num := d1 + d2

		total += util.MustParseInt(num)
	}

	return util.NewSolution(1, total)
}

func (d Day01) Part2(input string) *util.Solution {
	values := util.Lines(input)

	total := 0

	for _, value := range values {
		d1 := ""
		d2 := ""

		chars := strings.Split(value, "")

		for i, c := range chars {
			if i+2 < len(value) {
				potential3Num, has3 := TEXT_TO_DIGIT[c+chars[i+1]+chars[i+2]]

				if has3 {
					if d1 == "" {
						d1 = potential3Num
						continue
					}

					d2 = potential3Num
					continue
				}
			}

			if i+3 < len(value) {
				potential4Num, has4 := TEXT_TO_DIGIT[c+chars[i+1]+chars[i+2]+chars[i+3]]
				if has4 {
					if d1 == "" {
						d1 = potential4Num
						continue
					}

					d2 = potential4Num
					continue
				}
			}

			if i+4 < len(value) {
				potential5Num, has5 := TEXT_TO_DIGIT[c+chars[i+1]+chars[i+2]+chars[i+3]+chars[i+4]]

				if has5 {
					if d1 == "" {
						d1 = potential5Num
						continue
					}

					d2 = potential5Num
					continue
				}
			}

			_, err := strconv.Atoi(c)

			if err == nil && d1 == "" {
				d1 = c
				continue
			} else if err == nil {
				d2 = c
				continue
			}
		}

		if d2 == "" {
			d2 = d1
		}

		num := d1 + d2

		total += util.MustParseInt(num)
	}

	return util.NewSolution(2, total)
}
