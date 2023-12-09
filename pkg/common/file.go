package common

import (
	"fmt"
	"os"
	"path"
)

func LoadDailyInput(day int) string {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error while fetching current dir", err)
		os.Exit(1)
	}

	dayNr := fmt.Sprintf("%02d", day)

	if len(dayNr) > 2 {
		dayNr = fmt.Sprint(day)
	}

	inputPath := path.Join(cwd, "data", fmt.Sprintf("%s.txt", dayNr))

	rawInput, err := os.ReadFile(inputPath)

	if err != nil {
		fmt.Println("Error while reading file", err)
		os.Exit(1)
	}

	return string(rawInput)
}
