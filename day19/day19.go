package day19

import (
	"regexp"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const ACCEPTED = "A"
const REJECTED = "R"

var PIPELINE_MATCHER = regexp.MustCompile(`(?P<label>\w+){(?P<instructions>([asmx]+[<>]{1}\d+:\w+,?)+)(?P<endBucket>\w+)}`)
var INSTRUCTION_MATCHER = regexp.MustCompile(`(?P<property>[asmx]{1})+(?P<operator>[<>]{1})(?P<value>\d+):(?P<targetBucket>\w+)`)
var ITEM_MATCHER = regexp.MustCompile(`{x=(?P<x>\d+),m=(?P<m>\d+),a=(?P<a>\d+),s=(?P<s>\d+)}`)

type Day19 struct{}

func (d Day19) Part1(input string) *common.Solution {
	groups := strings.Split(input, "\n\n")
	rawPipelines := util.Lines(groups[0])
	rawItems := util.Lines(groups[1])

	acceptedItems := make([]*Item, 0)
	rejectedItems := make([]*Item, 0)

	pipelines := make(map[string]*Pipeline)
	items := make([]*Item, len(rawItems))

	for _, rawPipeline := range rawPipelines {
		pipelineParts := util.MatchNamedSubgroups(PIPELINE_MATCHER, rawPipeline)
		pipeline := Pipeline{endBucket: pipelineParts["endBucket"], instructions: make([]Instruction, 0)}

		for _, rawInstruction := range strings.Split(pipelineParts["instructions"], ",") {
			if rawInstruction == "" {
				continue
			}

			instrParts := util.MatchNamedSubgroups(INSTRUCTION_MATCHER, rawInstruction)
			property := instrParts["property"]

			instruction := Instruction{
				property:     property,
				operator:     instrParts["operator"],
				value:        util.MustParseInt(instrParts["value"]),
				targetBucket: instrParts["targetBucket"],
			}

			pipeline.instructions = append(pipeline.instructions, instruction)
		}

		pipelines[pipelineParts["label"]] = &pipeline
	}

	for i, rawItem := range rawItems {
		properties := util.MatchNamedSubgroups(ITEM_MATCHER, rawItem)

		items[i] = &Item{
			x: util.MustParseInt(properties["x"]),
			m: util.MustParseInt(properties["m"]),
			a: util.MustParseInt(properties["a"]),
			s: util.MustParseInt(properties["s"]),
		}
	}

	for _, item := range items {
		current := pipelines["in"]

		for current != nil {
			nextKey := current.endBucket

			for _, instr := range current.instructions {
				if instr.accepts(item) {
					nextKey = instr.targetBucket
					break
				}
			}

			if nextKey == REJECTED {
				rejectedItems = append(rejectedItems, item)
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
		totalScore += item.x + item.m + item.a + item.s
	}

	return common.NewSolution(1, totalScore)
}

func (d Day19) Part2(input string) *common.Solution {
	return common.NewSolution(2, -1)
}

type Pipeline struct {
	instructions []Instruction
	endBucket    string
}

type Instruction struct {
	property     string
	operator     string
	value        int
	targetBucket string
}

func (instr Instruction) accepts(item *Item) bool {
	propertyValue := -1

	if instr.property == "x" {
		propertyValue = item.x
	} else if instr.property == "m" {
		propertyValue = item.m
	} else if instr.property == "a" {
		propertyValue = item.a
	} else if instr.property == "s" {
		propertyValue = item.s
	}

	if instr.operator == ">" {
		return propertyValue > instr.value
	}

	return propertyValue < instr.value
}

type Item struct {
	x int
	m int
	a int
	s int
}
