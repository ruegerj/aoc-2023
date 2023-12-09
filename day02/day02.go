package day02

import (
	"regexp"
	"strings"

	"github.com/ruegerj/aoc-2023/pkg/common"
	"github.com/ruegerj/aoc-2023/pkg/util"
)

type Day02 struct{}

func (d Day02) Part1(input string) *common.Solution {
	cubeBalance := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := parseGames(input)

	totalPossibleGames := 0

	for _, game := range games {
		if game.isPossible(cubeBalance) {
			totalPossibleGames += game.id
		}
	}

	return common.NewSolution(1, totalPossibleGames)
}

func (d Day02) Part2(input string) *common.Solution {
	games := parseGames(input)

	powerSum := 0

	for _, game := range games {
		powerSum += game.calcPower()
	}

	return common.NewSolution(2, powerSum)
}

func parseGames(input string) []Game {
	rawGames := util.Lines(input)
	games := make([]Game, 0)

	for _, rawGame := range rawGames {
		gameDataMatcher := regexp.MustCompile(`^Game (?P<nr>\d+): (?P<rounds>.*)$`)
		gameData := util.MatchNamedSubgroups(gameDataMatcher, rawGame)

		game := Game{id: util.MustParseInt(gameData["nr"]), maxRed: 0, maxBlue: 0, maxGreen: 0}

		for _, round := range strings.Split(gameData["rounds"], "; ") {
			steps := strings.Split(round, " ")

			for i := 0; i < len(steps); i += 2 {
				count := util.MustParseInt(steps[i])
				color := steps[i+1]

				if (color == "red" || color == "red,") && game.maxRed < count {
					game.maxRed = count
					continue
				}

				if (color == "blue" || color == "blue,") && game.maxBlue < count {
					game.maxBlue = count
					continue
				}

				if (color == "green" || color == "green,") && game.maxGreen < count {
					game.maxGreen = count
				}
			}
		}

		games = append(games, game)
	}

	return games
}

type Game struct {
	id       int
	maxRed   int
	maxBlue  int
	maxGreen int
}

func (game *Game) isPossible(cubeBalance map[string]int) bool {
	return game.maxRed <= cubeBalance["red"] &&
		game.maxGreen <= cubeBalance["green"] &&
		game.maxBlue <= cubeBalance["blue"]
}

func (game *Game) calcPower() int {
	return game.maxRed * game.maxGreen * game.maxBlue
}
