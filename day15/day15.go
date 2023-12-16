package day15

import (
	"errors"
	"regexp"
	"slices"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

const ADD_OR_UPDATE = "="
const REMOVE = "-"

var LENS_MATCHER = regexp.MustCompile(`(?P<label>\w+)(?P<operation>[-=]{1})(?P<power>\d?)`)

type Day15 struct{}

func (d Day15) Part1(input string) *common.Solution {
	totalHashSum := 0

	for _, part := range strings.Split(input, ",") {
		totalHashSum += hash(part)
	}

	return common.NewSolution(1, totalHashSum)
}

func (d Day15) Part2(input string) *common.Solution {
	hashMap := newHashMap()

	for _, part := range strings.Split(input, ",") {
		lens := newLens(part)

		if lens.operation == ADD_OR_UPDATE {
			hashMap.addOrUpdate(lens)
			continue
		}

		if lens.operation == REMOVE {
			hashMap.remove(lens)
		}
	}

	focusingPower := 0

	for box := 1; box <= 256; box++ {
		lenses := hashMap.get(box - 1)

		for pos := 1; pos <= len(lenses); pos++ {
			lense := lenses[pos-1]
			focusingPower += box * pos * lense.power
		}
	}

	return common.NewSolution(2, focusingPower)
}

func hash(value string) int {
	hash := 0

	for _, char := range value {
		hash += int(char)
		hash *= 17
		hash = hash % 256
	}

	return hash
}

type hashMap struct {
	internalMap map[int][]lens
}

func newHashMap() *hashMap {
	instance := &hashMap{
		internalMap: make(map[int][]lens, 256),
	}

	for i := 0; i < 256; i++ {
		instance.internalMap[i] = make([]lens, 0)
	}

	return instance
}

func (hm *hashMap) get(hash int) []lens {
	lenses, found := hm.internalMap[hash]

	if !found {
		panic(errors.New("couldn't find lenses for hash"))
	}

	return lenses
}

func (hm *hashMap) addOrUpdate(new lens) {
	lenses, _ := hm.internalMap[new.hashCode()]

	index := slices.IndexFunc(lenses, func(l lens) bool {
		return l.label == new.label
	})

	if index >= 0 {
		lenses[index] = new
		return
	}

	hm.internalMap[new.hashCode()] = append(lenses, new)
}

func (hm *hashMap) remove(old lens) {
	lenses, _ := hm.internalMap[old.hashCode()]

	index := slices.IndexFunc(lenses, func(l lens) bool {
		return l.label == old.label
	})

	if index < 0 {
		return
	}

	hm.internalMap[old.hashCode()] = util.RemoveIndex(index, lenses)
}

type lens struct {
	label     string
	operation string
	power     int
}

func newLens(raw string) lens {
	parts := util.MatchNamedSubgroups(LENS_MATCHER, raw)

	lens := lens{
		label:     parts["label"],
		operation: parts["operation"],
	}

	power, hasPower := parts["power"]

	if hasPower && power != "" {
		lens.power = util.MustParseInt(power)
	}

	return lens
}

func (l lens) hashCode() int {
	return hash(l.label)
}
