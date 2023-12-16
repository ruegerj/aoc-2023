package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ruegerj/aoc-2023/day01"
	"github.com/ruegerj/aoc-2023/day02"
	"github.com/ruegerj/aoc-2023/day03"
	"github.com/ruegerj/aoc-2023/day04"
	"github.com/ruegerj/aoc-2023/day05"
	"github.com/ruegerj/aoc-2023/day06"
	"github.com/ruegerj/aoc-2023/day07"
	"github.com/ruegerj/aoc-2023/day08"
	"github.com/ruegerj/aoc-2023/day09"
	"github.com/ruegerj/aoc-2023/day10"
	"github.com/ruegerj/aoc-2023/day11"
	"github.com/ruegerj/aoc-2023/day12"
	"github.com/ruegerj/aoc-2023/day13"
	"github.com/ruegerj/aoc-2023/day14"
	"github.com/ruegerj/aoc-2023/day15"
	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

func main() {
	fmt.Println(`          ____   _____   ___   ___ ___  ____  `)
	fmt.Println(`    /\   / __ \ / ____| |__ \ / _ \__ \|___ \ `)
	fmt.Println(`   /  \ | |  | | |         ) | | | | ) | __) |`)
	fmt.Println(`  / /\ \| |  | | |        / /| | | |/ / |__ < `)
	fmt.Println(` / ____ \ |__| | |____   / /_| |_| / /_ ___) |`)
	fmt.Println(`/_/    \_\____/ \_____| |____|\___/____|____/ `)
	fmt.Println("----------------------------------------------")
	fmt.Println("🎄 Happy Coding & festive season")

	dayArg := os.Args[1]
	printHelp := strings.Contains(dayArg, "help") || strings.Contains(dayArg, "h")

	if printHelp {
		fmt.Println("usage: go run . <day-nr>")
		return
	}

	dayNr, err := strconv.Atoi(dayArg)

	if err != nil {
		fmt.Println("❌ Invalid day number...")
		return
	}

	dayRegistry := map[int]common.Day{
		1:  day01.Day01{},
		2:  day02.Day02{},
		3:  day03.Day03{},
		4:  day04.Day04{},
		5:  day05.Day05{},
		6:  day06.Day06{},
		7:  day07.Day07{},
		8:  day08.Day08{},
		9:  day09.Day09{},
		10: day10.Day10{},
		11: day11.Day11{},
		12: day12.Day12{},
		13: day13.Day13{},
		14: day14.Day14{},
		15: day15.Day15{},
	}
	requestedDay := dayRegistry[dayNr]

	if requestedDay == nil {
		fmt.Println("🛠  Not implemented")
		return
	}

	runDay(dayNr, requestedDay)
}

func runDay(nr int, day common.Day) {
	input := common.LoadDailyInput(nr)
	normalizedNr := util.PadNumber(nr)

	fmt.Printf("\n⭐️ Day %s\n", normalizedNr)

	start1 := time.Now()
	solution1 := day.Part1(input)
	solution1.Print(time.Since(start1))

	start2 := time.Now()
	solution2 := day.Part2(input)
	solution2.Print(time.Since(start2))
}
