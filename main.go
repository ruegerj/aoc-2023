package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ruegerj/aoc-2023/util"
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

	dayRegistry := map[int]util.Day{}
	requestedDay := dayRegistry[dayNr]

	if requestedDay == nil {
		fmt.Println("🛠  Not implemented")
		return
	}

	runDay(dayNr, requestedDay)
}

func runDay(nr int, day util.Day) {
	input := util.LoadDailyInput(nr)
	normalizedNr := util.PadNumber(nr)

	fmt.Printf("\n⭐️ Day %s\n", normalizedNr)

	start1 := time.Now()
	solution1 := day.Part1(input)
	solution1.Print(time.Since(start1))

	start2 := time.Now()
	solution2 := day.Part2(input)
	solution2.Print(time.Since(start2))
}
