package day17

import (
	"container/heap"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day17 struct{}

func (d Day17) Part1(input string) *common.Solution {
	graph := parseGraph(input)
	startVector := &Vector{x: 0, y: 0, dx: 0, dy: 0, length: 0, priority: 0}

	pq := PriorityQueue{startVector}
	heap.Init(&pq)

	seen := mapset.NewSet[Visit]()

	minCost := -1

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Vector)
		visit := current.visit()

		if current.isEnd(graph) {
			minCost = current.priority
			break
		}

		if seen.Contains(visit) {
			continue
		}

		seen.Add(visit)

		sameDirection := &Vector{
			x:      current.x + current.dx,
			y:      current.y + current.dy,
			dy:     current.dy,
			dx:     current.dx,
			length: current.length + 1,
		}

		if current.length < 3 && !current.isStale() && sameDirection.isInBounds(graph) {
			sameDirection.priority = current.priority + graph[sameDirection.y][sameDirection.x]
			heap.Push(&pq, sameDirection)
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

			direction.y = current.y + direction.dy
			direction.x = current.x + direction.dx

			if direction.isInBounds(graph) {
				direction.priority = current.priority + graph[direction.y][direction.x]

				heap.Push(&pq, direction)
			}
		}
	}

	return common.NewSolution(1, minCost)
}

func (d Day17) Part2(input string) *common.Solution {
	graph := parseGraph(input)
	startVector := &Vector{x: 0, y: 0, dx: 0, dy: 0, length: 0, priority: 0}

	pq := PriorityQueue{startVector}
	heap.Init(&pq)

	seen := mapset.NewSet[Visit]()

	minCost := -1

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Vector)
		visit := current.visit()

		if current.isEnd(graph) && current.length >= 4 {
			minCost = current.priority
			break
		}

		if seen.Contains(visit) {
			continue
		}

		seen.Add(visit)

		sameDirection := &Vector{
			x:      current.x + current.dx,
			y:      current.y + current.dy,
			dy:     current.dy,
			dx:     current.dx,
			length: current.length + 1,
		}

		if current.length < 10 && !current.isStale() && sameDirection.isInBounds(graph) {
			sameDirection.priority = current.priority + graph[sameDirection.y][sameDirection.x]
			heap.Push(&pq, sameDirection)
		}

		if current.length < 4 && !current.isStale() {
			continue
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

			direction.y = current.y + direction.dy
			direction.x = current.x + direction.dx

			if direction.isInBounds(graph) {
				direction.priority = current.priority + graph[direction.y][direction.x]

				heap.Push(&pq, direction)
			}
		}
	}

	return common.NewSolution(2, minCost)
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

func (v *Vector) isStale() bool {
	return v.dx == 0 && v.dy == 0
}

func (v *Vector) isInBounds(graph [][]int) bool {
	return v.x >= 0 && v.x < len(graph[0]) && v.y >= 0 && v.y < len(graph)
}

func (v *Vector) isEnd(graph [][]int) bool {
	return v.x == len(graph[0])-1 && v.y == len(graph)-1
}

func (v *Vector) visit() Visit {
	return Visit{
		x:     v.x,
		y:     v.y,
		dx:    v.dx,
		dy:    v.dy,
		steps: v.length,
	}
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
