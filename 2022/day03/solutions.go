package solutions

import (
	"slices"
	"strings"
)

func Part1(input string) (prioritySum int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		firstHalf, secondHalf := splitStringInHalf(line)
		intersection := getIntersection(firstHalf, secondHalf)
		for _, letter := range intersection {
			prioritySum += getPriorityValue(letter)
		}
	}

	return
}

func splitStringInHalf(toSplit string) (firstHalf, secondHalf string) {
	length := len(toSplit)
	firstHalf = toSplit[:length/2]
	secondHalf = toSplit[length/2:]

	return
}

func getIntersection(firstHalf string, secondHalf string) (inter []byte) {
	string1 := firstHalf
	string2 := secondHalf
	if len(string1) > len(string2) {
		string1, string2 = string2, string1
	}

	for _, letter := range string1 {
		if !strings.Contains(string2, string(letter)) {
			continue
		}

		letterByte := byte(letter)
		if slices.Contains(inter, letterByte) {
			continue
		}

		inter = append(inter, letterByte)
	}

	return
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
