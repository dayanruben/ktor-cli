package utils

import (
	"slices"
	"strings"
)

// CleanProjectName returns project name stripped until the first forbidden symbol
func CleanProjectName(name string) string {
	forbiddenChars := map[rune]struct{}{'/': {}, '\\': {}, ':': {}, '<': {}, '>': {}, '"': {}, '?': {}, '*': {}, '|': {}}
	lastIndex := -1

	for i, ch := range name {
		if _, ok := forbiddenChars[ch]; ok {
			lastIndex = i
			break
		}
	}

	if lastIndex == -1 {
		return name
	}

	return name[:lastIndex]
}

func GetPackage(website string) string {
	parts := strings.Split(website, ".")
	slices.Reverse(parts)
	return strings.Join(parts, ".")
}
