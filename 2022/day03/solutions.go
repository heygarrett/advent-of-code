package solutions

import (
	"strings"
)

func Part1(input string) (prioritySum int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		strSlice := splitStringInHalf(line)
		intersection := getIntersection(strSlice)
		for _, letter := range intersection {
			prioritySum += getPriorityValue(byte(letter))
		}
	}

	return
}

func Part2(input string) (prioritySum int) {
	lines := strings.Split(input, "\n")

	group := []string{}
	for _, str := range lines {
		group = append(group, str)
		if len(group) == 3 {
			intersection := getIntersection(group)
			for _, letter := range intersection {
				prioritySum += getPriorityValue(byte(letter))
			}
			group = nil
		}
	}

	return
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

	newStrSlice := append([]string{set}, strSlice[2:]...)

	return getIntersection(newStrSlice)
}

func getPriorityValue(letter byte) (value int) {
	lowercaseOffset := 96
	capitalOffset := 38

	letterInt := int(letter)
	if letterInt > lowercaseOffset {
		value = letterInt - lowercaseOffset
	} else {
		value = letterInt - capitalOffset
	}

	return
}
