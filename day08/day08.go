package day08

import (
	"regexp"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day08 struct{}

const LEFT = "L"
const RIGHT = "R"

var NODE_MATCHER = regexp.MustCompile(`(?P<name>[A-Z0-9]{3}) = \((?P<left>[A-Z0-9]{3}), (?P<right>[A-Z0-9]{3})\)`)

func (d Day08) Part1(input string) *common.Solution {
	const START = "AAA"
	const END = "ZZZ"

	instructions, nodes := parseNetwork(input)

	steps := 0
	current, _ := nodes[START]

	for true {
		for _, direction := range instructions {
			if current.name == END {
				break
			}

			if direction == LEFT {
				current = current.left
			}

			if direction == RIGHT {
				current = current.right
			}

			steps++
		}

		if current.name == END {
			break
		}
	}

	return common.NewSolution(1, steps)
}

func (d Day08) Part2(input string) *common.Solution {
	instructions, nodes := parseNetwork(input)

	currentNodes := make([]*Node, 0)

	for _, node := range nodes {
		if node.name[2] == byte('A') {
			currentNodes = append(currentNodes, node)
		}
	}

	stepsToEnd := make([]int, len(currentNodes))

	for i, current := range currentNodes {
		steps := 0

		for true {
			for _, direction := range instructions {
				if current.name[2] == byte('Z') {
					break
				}

				if direction == LEFT {
					current = current.left
				}

				if direction == RIGHT {
					current = current.right
				}

				steps++
			}

			if current.name[2] == byte('Z') {
				break
			}
		}

		stepsToEnd[i] = steps
	}

	minStepsToReachAllEnds := util.LCM(stepsToEnd[0], stepsToEnd[1], stepsToEnd[2:]...)

	return common.NewSolution(2, minStepsToReachAllEnds)
}

func parseNetwork(input string) ([]string, map[string]*Node) {
	parts := strings.Split(input, "\n\n")
	instructions := strings.Split(parts[0], "")
	nodeDefinitions := util.Lines(parts[1])

	nodes := map[string]*Node{}

	for _, definition := range nodeDefinitions {
		matches := util.MatchNamedSubgroups(NODE_MATCHER, definition)
		name := matches["name"]
		left := matches["left"]
		right := matches["right"]

		node, nodeKnown := nodes[name]

		if !nodeKnown {
			node = &Node{name: name}
		}

		leftNode, leftKnown := nodes[left]

		if !leftKnown {
			leftNode = &Node{name: left}
			nodes[left] = leftNode
		}

		node.left = leftNode

		rightNode, rightKnown := nodes[right]

		if !rightKnown {
			rightNode = &Node{name: right}
			nodes[right] = rightNode
		}

		node.right = rightNode
		nodes[name] = node
	}

	return instructions, nodes
}

type Node struct {
	name  string
	left  *Node
	right *Node
}
