package day17

import (
	"container/heap"
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day17 struct{}

func (d Day17) Part1(input string) *common.Solution {
	matrix := parseGraph(input)

	startVector := &Vector{
		x:        0,
		y:        0,
		dx:       0,
		dy:       0,
		length:   0,
		priority: 0,
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, startVector)

	seen := mapset.NewSet[Visit]()

	minCost := 0

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Vector)
		visit := Visit{
			x:     current.x,
			y:     current.y,
			dx:    current.dx,
			dy:    current.dy,
			steps: current.length,
		}

		if current.x == len(matrix[0])-1 && current.y == len(matrix)-1 {
			minCost = current.priority
			break
		}

		if seen.Contains(visit) {
			continue
		}

		seen.Add(visit)

		nextSameDirection := &Vector{dy: current.dy, dx: current.dx, length: current.length + 1}

		if current.length < 3 && (current.dy != 0 || current.dx != 0) {
			nextY := current.y + nextSameDirection.dy
			nextX := current.x + nextSameDirection.dx

			if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
				nextSameDirection.x = nextX
				nextSameDirection.y = nextY
				nextSameDirection.priority = current.priority + matrix[nextY][nextX]

				heap.Push(&pq, nextSameDirection)
			}
		}

		directions := []*Vector{
			{dy: 0, dx: 1, length: 1},
			{dy: 1, dx: 0, length: 1},
			{dy: 0, dx: -1, length: 1},
			{dy: -1, dx: 0, length: 1},
		}

		for _, direction := range directions {
			isSame := direction.dx == current.dx && direction.dy == current.dy
			isInverse := direction.dx == -current.dx && direction.dy == -current.dy

			if isSame || isInverse {
				continue
			}

			nextY := current.y + direction.dy
			nextX := current.x + direction.dx

			if nextX >= 0 && nextX < len(matrix[0]) && nextY >= 0 && nextY < len(matrix) {
				direction.x = nextX
				direction.y = nextY
				direction.priority = current.priority + matrix[nextY][nextX]

				heap.Push(&pq, direction)
			}
		}
	}

	return common.NewSolution(1, minCost)
}

func (d Day17) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

func colored(value string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, value)
}

func parseGraph(input string) [][]int {
	lines := util.Lines(input)
	matrix := make([][]int, len(lines))

	for y, line := range lines {
		fields := strings.Split(line, "")
		matrix[y] = make([]int, len(fields))

		for x, field := range fields {
			matrix[y][x] = util.MustParseInt(field)
		}
	}

	return matrix
}

type Visit struct {
	x     int
	y     int
	dx    int
	dy    int
	steps int
}

type Vector struct {
	x      int
	y      int
	dx     int
	dy     int
	length int

	priority int
	index    int
}

// Array based priority queue which implements the container/heap interface
// taken from: https://pkg.go.dev/container/heap#example-package-PriorityQueue
type PriorityQueue []*Vector

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = pq[j].index
	pq[j].index = pq[i].index
}

func (pq *PriorityQueue) Push(item any) {
	length := len(*pq)
	newItem := item.(*Vector)
	newItem.index = length
	*pq = append(*pq, newItem)
}

func (pq *PriorityQueue) Pop() any {
	oldQueue := *pq
	length := len(oldQueue)
	popped := oldQueue[length-1]
	oldQueue[length-1] = nil // prevent memory leak
	popped.index = -1
	*pq = oldQueue[0 : length-1]
	return popped
}
