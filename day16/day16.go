package day16

import (
	"errors"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
	"golang.org/x/exp/maps"
)

const (
	LEFT_RIGHT = iota
	RIGHT_LEFT
	TOP_BOTTOM
	BOTTOM_UP
)

const EMPTY = "."
const LEFT_MIRROR = "/"
const RIGHT_MIRROR = "\\"
const VERTICAL_SPLITTER = "|"
const HORIZONTAL_SPLITTER = "-"

type Day16 struct{}

func (d Day16) Part1(input string) *common.Solution {
	tiles := util.Matrix(input, "")

	startTile := tile{y: 0, x: 0}
	visited := map[tile]int{}

	flow(startTile, LEFT_RIGHT, tiles, visited)

	energizedTiles := 0
	visitedTiles := mapset.NewSet[tile]()

	for _, key := range maps.Keys(visited) {
		if visitedTiles.Contains(key) {
			continue
		}

		energizedTiles++
		visitedTiles.Add(key)
	}

	return common.NewSolution(1, energizedTiles)
}

func (d Day16) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

func flow(current tile, direction int, tiles [][]string, visited map[tile]int) {
	if current.outOfBounds(tiles) {
		return
	}

	visitedFrom, hasVisit := visited[current]

	if hasVisit && visitedFrom == direction {
		return
	}

	item := tiles[current.y][current.x]

	visited[current] = direction

	switch item {
	case EMPTY:
		next := nextTileInDirection(current, direction)
		flow(next, direction, tiles, visited)
	case VERTICAL_SPLITTER:
		if direction == TOP_BOTTOM || direction == BOTTOM_UP {
			next := nextTileInDirection(current, direction)
			flow(next, direction, tiles, visited)
			return
		}

		flow(nextTileUp(current), BOTTOM_UP, tiles, visited)
		flow(nextTileDown(current), TOP_BOTTOM, tiles, visited)
	case HORIZONTAL_SPLITTER:
		if direction == LEFT_RIGHT || direction == RIGHT_LEFT {
			next := nextTileInDirection(current, direction)
			flow(next, direction, tiles, visited)
			return
		}

		flow(nextTileLeft(current), RIGHT_LEFT, tiles, visited)
		flow(nextTileRight(current), LEFT_RIGHT, tiles, visited)
	case LEFT_MIRROR:
		var next tile
		var newDirection int

		if direction == LEFT_RIGHT {
			next = nextTileUp(current)
			newDirection = BOTTOM_UP
		} else if direction == RIGHT_LEFT {
			next = nextTileDown(current)
			newDirection = TOP_BOTTOM
		} else if direction == TOP_BOTTOM {
			next = nextTileLeft(current)
			newDirection = RIGHT_LEFT
		} else if direction == BOTTOM_UP {
			next = nextTileRight(current)
			newDirection = LEFT_RIGHT
		} else {
			// TODO: remove
			fmt.Println(current, direction)
			panic(errors.New("unknown direction"))
		}

		flow(next, newDirection, tiles, visited)

	case RIGHT_MIRROR:
		var next tile
		var newDirection int

		if direction == LEFT_RIGHT {
			next = nextTileDown(current)
			newDirection = TOP_BOTTOM
		} else if direction == RIGHT_LEFT {
			next = nextTileUp(current)
			newDirection = BOTTOM_UP
		} else if direction == TOP_BOTTOM {
			next = nextTileRight(current)
			newDirection = LEFT_RIGHT
		} else if direction == BOTTOM_UP {
			next = nextTileLeft(current)
			newDirection = RIGHT_LEFT
		} else {
			// TODO: remove
			fmt.Println(current, direction)
			panic(errors.New("unknown direction"))
		}

		flow(next, newDirection, tiles, visited)
	}
}

func nextTileInDirection(current tile, direction int) tile {
	var next tile
	if direction == LEFT_RIGHT {
		next = tile{y: current.y, x: current.x + 1}
	} else if direction == RIGHT_LEFT {
		next = tile{y: current.y, x: current.x - 1}
	} else if direction == TOP_BOTTOM {
		next = tile{y: current.y + 1, x: current.x}
	} else if direction == BOTTOM_UP {
		next = tile{y: current.y - 1, x: current.x}
	} else {
		// TODO: remove
		fmt.Println(current, direction)
		panic(errors.New("unknown direction"))
	}

	return next
}

func nextTileLeft(current tile) tile {
	return tile{y: current.y, x: current.x - 1}
}

func nextTileRight(current tile) tile {
	return tile{y: current.y, x: current.x + 1}
}

func nextTileUp(current tile) tile {
	return tile{y: current.y - 1, x: current.x}
}

func nextTileDown(current tile) tile {
	return tile{y: current.y + 1, x: current.x}
}

type tile struct {
	y int
	x int
}

func (t tile) outOfBounds(matrix [][]string) bool {
	if t.x < 0 || t.y < 0 {
		return true
	}

	if t.x >= len(matrix[0]) {
		return true
	}

	if t.y >= len(matrix) {
		return true
	}

	return false
}
