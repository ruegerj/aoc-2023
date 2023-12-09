package util

import "regexp"

func MatchNamedSubgroups(exp *regexp.Regexp, input string) map[string]string {
	match := exp.FindStringSubmatch(input)

	subgroupValues := make(map[string]string)
	for i, name := range exp.SubexpNames() {
		if i > 0 && i <= len(match) {
			subgroupValues[name] = match[i]
		}
	}

	return subgroupValues
}
