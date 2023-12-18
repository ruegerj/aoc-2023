package day18

import (
	"regexp"
	"strconv"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

var INSTRUCTION_MATCHER = regexp.MustCompile(`(?P<direction>[A-Z]{1}) (?P<steps>\d+) \(#(?P<color>[a-z0-9]{6})\)`)

type Day18 struct{}

func (d Day18) Part1(input string) *common.Solution {
	points := make([]*Point, 0)

	var trenchLength int64 = 0
	last := &Point{y: 0, x: 0}

	for _, line := range util.Lines(input) {
		instruction := util.MatchNamedSubgroups(INSTRUCTION_MATCHER, line)
		direction := instruction["direction"]
		steps := util.MustParseInt64(instruction["steps"])

		trenchLength += steps
		point := getNextPoint(direction, steps, last)

		last = point
		points = append(points, point)
	}

	lagoonArea := shoelace(points, trenchLength)

	return common.NewSolution(1, lagoonArea)
}

func (d Day18) Part2(input string) *common.Solution {
	points := make([]*Point, 0)

	var trenchLength int64 = 0
	last := &Point{y: 0, x: 0}

	for _, line := range util.Lines(input) {
		instruction := util.MatchNamedSubgroups(INSTRUCTION_MATCHER, line)
		stepsHex := instruction["color"][0:5]
		direction := instruction["color"][5:6]
		steps, _ := strconv.ParseInt(stepsHex, 16, 64)

		trenchLength += steps
		point := getNextPoint(direction, steps, last)

		last = point
		points = append(points, point)
	}

	lagoonArea := shoelace(points, trenchLength)

	return common.NewSolution(2, lagoonArea)
}

func shoelace(points []*Point, borderLength int64) int64 {
	var doubledArea int64 = 0

	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		p2 := points[i+1]

		doubledArea += (p1.x * p2.y) - (p1.y * p2.x)
	}

	return util.Abs64(doubledArea)/2 + borderLength/2 + 1
}

func getNextPoint(direction string, steps int64, last *Point) *Point {
	var point *Point

	if direction == "R" || direction == "0" {
		point = &Point{y: last.y, x: last.x + steps}
	} else if direction == "L" || direction == "2" {
		point = &Point{y: last.y, x: last.x - steps}
	} else if direction == "U" || direction == "3" {
		point = &Point{y: last.y - steps, x: last.x}
	} else if direction == "D" || direction == "1" {
		point = &Point{y: last.y + steps, x: last.x}
	}

	return point
}

type Point struct {
	y int64
	x int64
}
