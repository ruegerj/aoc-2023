package day10

import (
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const START = "S"
const NORTH_SOUTH = "|"
const EAST_WEST = "-"
const NORTH_EAST = "L"
const NORTH_WEST = "J"
const SOUTH_WEST = "7"
const SOUTH_EAST = "F"
const GROUND = "."

type Day10 struct{}

func (d Day10) Part1(input string) *common.Solution {
	rows := util.Lines(input)

	pipes := make([][]*Pipe, len(rows))

	var startPipe *Pipe

	for y, row := range rows {
		rawPipes := strings.Split(row, "")
		pipes[y] = make([]*Pipe, len(rawPipes))

		for x, rawPipe := range rawPipes {
			pipe := &Pipe{
				x:    x,
				y:    y,
				kind: rawPipe,
			}

			if rawPipe == START {
				startPipe = pipe
			}

			pipes[y][x] = pipe
		}
	}

	for y, row := range pipes {
		for x, pipe := range row {
			switch pipe.kind {
			case NORTH_SOUTH:
				if inBounds(x, y-1, pipes) {
					pipe.connected = append(pipe.connected, pipes[y-1][x])
				}

				if inBounds(x, y+1, pipes) {
					pipe.connected = append(pipe.connected, pipes[y+1][x])
				}
			case EAST_WEST:
				if inBounds(x-1, y, pipes) {
					pipe.connected = append(pipe.connected, pipes[y][x-1])
				}

				if inBounds(x+1, y, pipes) {
					pipe.connected = append(pipe.connected, pipes[y][x+1])
				}
			case NORTH_EAST:
				if inBounds(x, y-1, pipes) {
					pipe.connected = append(pipe.connected, pipes[y-1][x])
				}

				if inBounds(x+1, y, pipes) {
					pipe.connected = append(pipe.connected, pipes[y][x+1])
				}
			case NORTH_WEST:
				if inBounds(x, y-1, pipes) {
					pipe.connected = append(pipe.connected, pipes[y-1][x])
				}

				if inBounds(x-1, y, pipes) {
					pipe.connected = append(pipe.connected, pipes[y][x-1])
				}
			case SOUTH_EAST:
				if inBounds(x, y+1, pipes) {
					pipe.connected = append(pipe.connected, pipes[y+1][x])
				}

				if inBounds(x+1, y, pipes) {
					pipe.connected = append(pipe.connected, pipes[y][x+1])

				}
			case SOUTH_WEST:
				if inBounds(x, y+1, pipes) {
					pipe.connected = append(pipe.connected, pipes[y+1][x])
				}

				if inBounds(x-1, y, pipes) {
					pipe.connected = append(pipe.connected, pipes[y][x-1])
				}
			}
		}
	}

	flatPipes := util.Flat(pipes)

	var current *Pipe

	for _, pipe := range flatPipes {
		if util.FirstOrDefault(pipe.connected, func(p *Pipe) bool { return p == startPipe }) != nil {
			current = pipe
			break
		}
	}

	path := []*Pipe{startPipe}

	for current != startPipe {
		prev := util.LastElement(path)
		path = append(path, current)
		current = util.FirstOrDefault(current.connected, func(p *Pipe) bool { return p != prev })
	}

	maxSteps := len(path) / 2

	return common.NewSolution(1, maxSteps)
}

func (d Day10) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

func inBounds(x, y int, matrix [][]*Pipe) bool {
	if x < 0 || x >= len(matrix[0]) {
		return false
	}

	if y < 0 || y >= len(matrix) {
		return false
	}

	return true
}

type Pipe struct {
	x         int
	y         int
	kind      string
	connected []*Pipe
}
