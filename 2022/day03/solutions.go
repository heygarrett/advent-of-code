package solutions

import (
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	prioritySum := 0
	for _, line := range lines {
		strSlice := splitStringInHalf(line)
		intersection := getIntersection(strSlice)
		for _, letter := range intersection {
			prioritySum += getPriorityValue(letter)
		}
	}

	return prioritySum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	prioritySum := 0
	group := []string{}
	for _, str := range lines {
		group = append(group, str)
		if len(group) == 3 {
			intersection := getIntersection(group)
			for _, letter := range intersection {
				prioritySum += getPriorityValue(letter)
			}
			group = nil
		}
	}

	return prioritySum
}

func splitStringInHalf(toSplit string) []string {
	length := len(toSplit)
	firstHalf := toSplit[:length/2]
	secondHalf := toSplit[length/2:]

	return []string{firstHalf, secondHalf}
}

func getIntersection(strSlice []string) string {
	if strSlice == nil {
		return ""
	}
	if len(strSlice) == 1 {
		return strSlice[0]
	}

	set := ""
	for _, letter := range strSlice[0] {
		if strings.Contains(set, string(letter)) {
			continue
		}
		if strings.Contains(strSlice[1], string(letter)) {
			set = set + string(letter)
		}
	}

	// exit early if there's no intersection
	if set == "" {
		return set
	}

	newStrSlice := append([]string{set}, strSlice[2:]...)

	return getIntersection(newStrSlice)
}

func getPriorityValue(letter rune) int {
	lowercaseOffset := 96
	capitalOffset := 38

	value := 0
	letterInt := int(letter)
	if letterInt > lowercaseOffset {
		value = letterInt - lowercaseOffset
	} else {
		value = letterInt - capitalOffset
	}

	return value
}
