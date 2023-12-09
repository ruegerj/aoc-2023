package common

import (
	"fmt"
	"time"
)

type Solution struct {
	PartNr int
	Result any
}

func NewSolution(partNr int, result any) *Solution {
	return &Solution{
		PartNr: partNr,
		Result: result,
	}
}

func (solution *Solution) Print(elapsed time.Duration) {
	arrayResult, is2dStringArray := solution.Result.([][]string)

	if !is2dStringArray {
		fmt.Printf("Part %d: %s (%s)\n", solution.PartNr, fmt.Sprint(solution.Result), elapsed)
		return
	}

	fmt.Printf("Part %d: (%s)\n", solution.PartNr, elapsed)

	for _, item := range arrayResult {
		fmt.Println(item)
	}
}
