package util

import (
	"errors"
)

func KeyByValue[K comparable, V comparable](m map[K]V, value V) K {
	for k, v := range m {
		if v == value {
			return k
		}
	}

	panic(errors.New("Map doesn't contain a key which maps to the provided value"))
}
