package solutions

import (
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	prioritySum := 0
	for _, line := range lines {
		bisectedLine := splitStringInHalf(line)
		intersection := getIntersection(bisectedLine)
		for _, letter := range intersection {
			prioritySum += getPriorityValue(letter)
		}
	}

	return prioritySum
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	prioritySum := 0
	lineGroup := []string{}
	for _, str := range lines {
		lineGroup = append(lineGroup, str)
		if len(lineGroup) == 3 {
			intersection := getIntersection(lineGroup)
			for _, letter := range intersection {
				prioritySum += getPriorityValue(letter)
			}
			lineGroup = nil
		}
	}

	return prioritySum
}

func splitStringInHalf(toSplit string) []string {
	halfLength := len(toSplit) / 2
	firstHalf := toSplit[:halfLength]
	secondHalf := toSplit[halfLength:]

	return []string{firstHalf, secondHalf}
}

func getIntersection(listOfStrings []string) string {
	switch len(listOfStrings) {
	case 0:
		return ""
	case 1:
		return listOfStrings[0]
	}

	set := ""
	for _, letter := range listOfStrings[0] {
		if strings.Contains(set, string(letter)) {
			continue
		}
		if strings.Contains(listOfStrings[1], string(letter)) {
			set = set + string(letter)
		}
	}

	// exit early if there's no intersection
	if set == "" {
		return set
	}

	newListOfStrings := append([]string{set}, listOfStrings[2:]...)

	return getIntersection(newListOfStrings)
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
