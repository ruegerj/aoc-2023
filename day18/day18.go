package day18

import (
	"regexp"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const RIGHT = "R"
const LEFT = "L"
const UP = "U"
const DOWN = "D"

var INSTRUCTION_MATCHER = regexp.MustCompile(`(?P<direction>[A-Z]{1}) (?P<steps>\d+) \((?P<color>#[a-z0-9]{6})\)`)

type Day18 struct{}

func (d Day18) Part1(input string) *common.Solution {
	points := make([]*Point, 0)

	trenchLength := 0
	last := &Point{y: 0, x: 0}

	for _, line := range util.Lines(input) {
		instruction := util.MatchNamedSubgroups(INSTRUCTION_MATCHER, line)
		direction := instruction["direction"]
		steps := util.MustParseInt(instruction["steps"])

		var point *Point

		if direction == RIGHT {
			point = &Point{y: last.y, x: last.x + steps}
		} else if direction == LEFT {
			point = &Point{y: last.y, x: last.x - steps}
		} else if direction == UP {
			point = &Point{y: last.y - steps, x: last.x}
		} else if direction == DOWN {
			point = &Point{y: last.y + steps, x: last.x}
		}

		trenchLength += steps
		last = point
		points = append(points, point)
	}

	doubledArea := 0

	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		p2 := points[i+1]

		doubledArea += (p1.x * p2.y) - (p1.y * p2.x)
	}

	lagoonArea := util.Abs(doubledArea)/2 + trenchLength/2 + 1

	return common.NewSolution(1, lagoonArea)
}

func (d Day18) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

type Point struct {
	y int
	x int
}
