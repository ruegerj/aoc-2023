package day10

import (
	"slices"
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
	pipes, startPipe := parsePipeSystem(input)
	flatPipes := util.Flat(pipes)

	current := util.FirstOrDefault(flatPipes, func(fp *Pipe) bool {
		connectsToStartPipe := func(p *Pipe) bool { return p == startPipe }
		return util.FirstOrDefault(fp.connected, connectsToStartPipe) != nil
	})

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
	pipes, startPipe := parsePipeSystem(input)
	flatPipes := util.Flat(pipes)

	current := util.FirstOrDefault(flatPipes, func(fp *Pipe) bool {
		connectsToStartPipe := func(p *Pipe) bool { return p == startPipe }
		return util.FirstOrDefault(fp.connected, connectsToStartPipe) != nil
	})

	startPipe.visited = true
	path := []*Pipe{startPipe}

	for current != startPipe {
		prev := util.LastElement(path)
		current.visited = true
		path = append(path, current)
		current = util.FirstOrDefault(current.connected, func(p *Pipe) bool { return p != prev })
	}

	// pad entire matrix with ground (.)
	paddedPipes := groundPadPipes(pipes)

	// remove all non-path pipes outside the loop by flood-fill
	topLeftGround := paddedPipes[0][0]
	floodVisit(topLeftGround, paddedPipes)

	includedTiles := 0
	flatPaddedPipes := util.Flat(paddedPipes)

	for _, pipe := range flatPaddedPipes {
		if pipe.visited {
			continue
		}

		pipe.visited = true
		currentY := pipe.y - 1
		crossCount := 0

		for inBounds(pipe.x, currentY, paddedPipes) {
			pipe := paddedPipes[currentY][pipe.x]
			if slices.Contains(path, pipe) {
				crossCount += calcCrossCount(pipe)
			}
			currentY--
		}

		if crossCount%2 != 0 {
			includedTiles++
		}
	}

	return common.NewSolution(2, includedTiles)
}

func calcCrossCount(pipe *Pipe) int {
	if pipe.kind == EAST_WEST || pipe.kind == NORTH_EAST || pipe.kind == SOUTH_EAST {
		return 1
	}

	return 0
}

func floodVisit(pipe *Pipe, pipes [][]*Pipe) {
	if pipe.visited {
		return
	}

	pipe.visited = true

	if inBounds(pipe.x, pipe.y-1, pipes) {
		floodVisit(pipes[pipe.y-1][pipe.x], pipes)
	}

	if inBounds(pipe.x+1, pipe.y, pipes) {
		floodVisit(pipes[pipe.y][pipe.x+1], pipes)
	}

	if inBounds(pipe.x, pipe.y+1, pipes) {
		floodVisit(pipes[pipe.y+1][pipe.x], pipes)
	}

	if inBounds(pipe.x-1, pipe.y, pipes) {
		floodVisit(pipes[pipe.y][pipe.x-1], pipes)
	}
}

func groundPadPipes(pipes [][]*Pipe) [][]*Pipe {
	height := len(pipes) + 2
	width := len(pipes[0]) + 2
	padded := make([][]*Pipe, height)

	for i := 0; i < len(padded); i++ {
		if i == 0 || i == len(padded)-1 {
			padded[i] = make([]*Pipe, width)

			for j := 0; j < width; j++ {
				padded[i][j] = newGround(j, i)
			}

			continue
		}

		padded[i] = []*Pipe{newGround(0, i)}

		for _, pipe := range pipes[i-1] {
			pipe.x = pipe.x + 1
			pipe.y = pipe.y + 1
			padded[i] = append(padded[i], pipe)
		}

		padded[i] = append(padded[i], newGround(width-1, i))
	}

	return padded
}

func parsePipeSystem(input string) ([][]*Pipe, *Pipe) {
	var startPipe *Pipe
	rows := util.Lines(input)
	pipes := make([][]*Pipe, len(rows))

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

	return pipes, startPipe
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
	visited   bool
}

func newGround(x, y int) *Pipe {
	return &Pipe{
		x:    x,
		y:    y,
		kind: GROUND,
	}
}
