package util

import "fmt"

func PadNumber(nr int) string {
	paddedNumber := fmt.Sprintf("%02d", nr)

	if len(paddedNumber) <= 2 {
		return paddedNumber
	}

	return fmt.Sprint(nr)
}
