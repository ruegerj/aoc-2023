package day08

import (
	"regexp"
	"strings"

	"github.com/ruegerj/aoc-2023/util"
)

type Day08 struct{}

const START = "AAA"
const END = "ZZZ"
const LEFT = "L"
const RIGHT = "R"

func (d Day08) Part1(input string) *util.Solution {
	parts := strings.Split(input, "\n\n")
	instructions := strings.Split(parts[0], "")
	nodeDefinitions := util.Lines(parts[1])

	NODE_MATCHER := regexp.MustCompile(`(?P<name>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)`)

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

	return util.NewSolution(1, steps)
}

func (d Day08) Part2(input string) *util.Solution {
	return util.NewSolution(2, -1)
}

type Node struct {
	name  string
	left  *Node
	right *Node
}
