package utils

import "regexp"

func ParseRegexFind(regex *regexp.Regexp, matchResult []string) map[string]string {
	parseMap := make(map[string]string)

	for i, name := range regex.SubexpNames() {
		if i > 0 && i <= len(matchResult) {
			parseMap[name] = matchResult[i]
		}
	}
	return parseMap
}
