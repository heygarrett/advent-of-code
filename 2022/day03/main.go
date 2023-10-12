package main

import (
	reader "2022"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	input, err := reader.ReadFileFromFS(os.DirFS("day03"), "input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1 := Part1(input)
	fmt.Printf("Part 1: %d\n", part1)
}

func Part1(input string) (prioritySum int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		lineLength := len(line)
		firstHalf := line[:lineLength/2]
		secondHalf := line[lineLength/2:]
		intersection := getIntersection(firstHalf, secondHalf)
		for _, letter := range intersection {
			prioritySum += getPriorityValue(letter)
		}
	}

	return
}

func getIntersection(firstHalf string, secondHalf string) (inter []byte) {
	for _, letter := range firstHalf {
		if strings.Contains(secondHalf, string(letter)) && !slices.Contains(inter, byte(letter)) {
			inter = append(inter, byte(letter))
		}
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
