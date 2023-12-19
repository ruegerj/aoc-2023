package day19

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
	"golang.org/x/exp/maps"
)

const ACCEPTED = "A"
const REJECTED = "R"
const GT = ">"
const LT = "<"
const START_PIPELINE = "in"

var PIPELINE_MATCHER = regexp.MustCompile(`(?P<label>\w+){(?P<instructions>([xmas]+[<>]{1}\d+:\w+,?)+),(?P<elsePipeline>\w+)}`)
var INSTRUCTION_MATCHER = regexp.MustCompile(`(?P<property>[xmas]{1})+(?P<operator>[<>]{1})(?P<value>\d+):(?P<targetPipeline>\w+)`)
var ITEM_MATCHER = regexp.MustCompile(`{x=(?P<x>\d+),m=(?P<m>\d+),a=(?P<a>\d+),s=(?P<s>\d+)}`)

type Day19 struct{}

func (d Day19) Part1(input string) *common.Solution {
	pipelines, items := parseSystem(input)

	acceptedItems := make([]Item, 0)

	for _, item := range items {
		current := pipelines[START_PIPELINE]

		for current != nil {
			nextKey := current.elsePipeline

			for _, instr := range current.instructions {
				if instr.accepts(item) {
					nextKey = instr.targetPipeline
					break
				}
			}

			if nextKey == REJECTED {
				break
			}

			if nextKey == ACCEPTED {
				acceptedItems = append(acceptedItems, item)
				break
			}

			current = pipelines[nextKey]
		}
	}

	totalScore := 0

	for _, item := range acceptedItems {
		totalScore += item["x"] + item["m"] + item["a"] + item["s"]
	}

	return common.NewSolution(1, totalScore)

}

func (d Day19) Part2(input string) *common.Solution {
	pipelines := flatParsePipelines(input)

	root := buildKdTree(pipelines[START_PIPELINE], pipelines)

	startItem := RangeItem{
		"x": Range{from: 1, to: 4000},
		"m": Range{from: 1, to: 4000},
		"a": Range{from: 1, to: 4000},
		"s": Range{from: 1, to: 4000},
	}

	foundRanges := root.walk(startItem)

	totalPossibilities := int64(0)

	for _, item := range foundRanges {
		xPossibilities := int64(item["x"].to - item["x"].from + 1)
		mPossibilities := int64(item["m"].to - item["m"].from + 1)
		aPossibilities := int64(item["a"].to - item["a"].from + 1)
		sPossibilities := int64(item["s"].to - item["s"].from + 1)

		totalPossibilities += xPossibilities * mPossibilities * aPossibilities * sPossibilities
	}

	return common.NewSolution(2, totalPossibilities)
}

func buildKdTree(current *Pipeline, pipelines map[string]*Pipeline) *KDNode {
	node := &KDNode{
		property:  current.property,
		operator:  current.operator,
		threshold: current.threshold,
	}

	leftKey := current.targetPipeline
	rightKey := current.elsePipeline

	if leftKey == ACCEPTED || leftKey == REJECTED {
		node.left = &KDNode{property: leftKey}
	}

	if rightKey == ACCEPTED || rightKey == REJECTED {
		node.right = &KDNode{property: rightKey}
	}

	if node.left == nil {
		node.left = buildKdTree(pipelines[leftKey], pipelines)
	}

	if node.right == nil {
		node.right = buildKdTree(pipelines[rightKey], pipelines)
	}

	return node
}

func parseSystem(input string) (map[string]*Pipeline, []Item) {
	groups := strings.Split(input, "\n\n")
	rawPipelines := util.Lines(groups[0])
	rawItems := util.Lines(groups[1])

	pipelines := make(map[string]*Pipeline)
	items := make([]Item, len(rawItems))

	for _, rawPipeline := range rawPipelines {
		pipelineParts := util.MatchNamedSubgroups(PIPELINE_MATCHER, rawPipeline)
		pipeline := Pipeline{elsePipeline: pipelineParts["elsePipeline"], instructions: make([]Instruction, 0)}

		for _, rawInstruction := range strings.Split(pipelineParts["instructions"], ",") {
			if rawInstruction == "" {
				continue
			}

			instrParts := util.MatchNamedSubgroups(INSTRUCTION_MATCHER, rawInstruction)
			property := instrParts["property"]

			instruction := Instruction{
				property:       property,
				operator:       instrParts["operator"],
				value:          util.MustParseInt(instrParts["value"]),
				targetPipeline: instrParts["targetPipeline"],
			}

			pipeline.instructions = append(pipeline.instructions, instruction)
		}

		pipelines[pipelineParts["label"]] = &pipeline
	}

	for i, rawItem := range rawItems {
		properties := util.MatchNamedSubgroups(ITEM_MATCHER, rawItem)

		items[i] = map[string]int{
			"x": util.MustParseInt(properties["x"]),
			"m": util.MustParseInt(properties["m"]),
			"a": util.MustParseInt(properties["a"]),
			"s": util.MustParseInt(properties["s"]),
		}
	}

	return pipelines, items
}

func flatParsePipelines(input string) map[string]*Pipeline {
	groups := strings.Split(input, "\n\n")
	rawPipelines := util.Lines(groups[0])

	pipelines := make(map[string]*Pipeline)

	for _, rawPipeline := range rawPipelines {
		pipelineParts := util.MatchNamedSubgroups(PIPELINE_MATCHER, rawPipeline)
		rawInstructions := strings.Split(pipelineParts["instructions"], ",")

		baseLabel := pipelineParts["label"]
		for i, rawInstruction := range rawInstructions {
			instrParts := util.MatchNamedSubgroups(INSTRUCTION_MATCHER, rawInstruction)

			pipeline := &Pipeline{
				property:       instrParts["property"],
				operator:       instrParts["operator"],
				threshold:      util.MustParseInt(instrParts["value"]),
				targetPipeline: instrParts["targetPipeline"],
			}

			label := baseLabel

			if i > 0 {
				label += fmt.Sprint(i + 1)
			}

			if i+1 < len(rawInstructions) {
				pipeline.elsePipeline = baseLabel + fmt.Sprint(i+2)
				pipelines[label] = pipeline
				continue
			}

			pipeline.elsePipeline = pipelineParts["elsePipeline"]
			pipelines[label] = pipeline
		}
	}

	return pipelines
}

type KDNode struct {
	property  string
	operator  string
	threshold int
	left      *KDNode
	right     *KDNode
}

func (n *KDNode) walk(item RangeItem) []RangeItem {
	if n.property == ACCEPTED {
		return []RangeItem{item}
	}

	if n.property == REJECTED {
		return []RangeItem{}
	}

	var left, right Range
	propertyRange := item[n.property]

	if n.operator == GT {
		left = Range{from: n.threshold + 1, to: propertyRange.to}
		right = Range{from: propertyRange.from, to: n.threshold}
	} else if n.operator == LT {
		left = Range{from: propertyRange.from, to: n.threshold - 1}
		right = Range{from: n.threshold, to: propertyRange.to}
	}

	leftCopy := item.DeepCopy()
	leftCopy[n.property] = left
	rightCopy := item.DeepCopy()
	rightCopy[n.property] = right

	leftNodes := n.left.walk(leftCopy)
	rightNodes := n.right.walk(rightCopy)

	return append(leftNodes, rightNodes...)
}

type RangeItem map[string]Range

func (item RangeItem) DeepCopy() RangeItem {
	copy := make(RangeItem, len(item))

	for _, key := range maps.Keys(item) {
		originalRange := item[key]
		rangeCopy := Range{from: originalRange.from, to: originalRange.to}

		copy[key] = rangeCopy
	}

	return copy
}

type Range struct {
	from int
	to   int
}

type Item map[string]int

type Pipeline struct {
	property       string
	operator       string
	threshold      int
	targetPipeline string
	elsePipeline   string
	// part 1 only
	instructions []Instruction
}

type Instruction struct {
	property       string
	operator       string
	value          int
	targetPipeline string
}

func (instr Instruction) accepts(item Item) bool {
	propertyValue := item[instr.property]

	if instr.operator == ">" {
		return propertyValue > instr.value
	}

	return propertyValue < instr.value
}
