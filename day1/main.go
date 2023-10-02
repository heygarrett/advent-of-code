package main

import (
	reader "2022"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := reader.NewInputFromFS(os.DirFS("day1"), "input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	answer, nil := GetMaxCalories(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(answer)
}

func GetMaxCalories(input string) (int, error) {
	parsedInput := parseInput(input)
	elfSums := make([]int, 0)
	for _, elf := range parsedInput {
		intList, err := convertStrListToIntList(elf)
		if err != nil {
			return 0, err
		}
		intTotal := sumAll(intList)
		elfSums = append(elfSums, intTotal)
	}

	return slices.Max(elfSums), nil
}

func parseInput(input string) [][]string {
	splitOnBlankLines := strings.Split(input, "\n\n")
	splitOnNewLines := make([][]string, 0)
	for _, contiguousLines := range splitOnBlankLines {
		splitOnNewLines = append(
			splitOnNewLines,
			strings.Split(contiguousLines, "\n"),
		)
	}

	return splitOnNewLines
}

func sumAll(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

func convertStrListToIntList(strs []string) ([]int, error) {
	result := make([]int, 0)
	for _, s := range strs {
		converted, err := strconv.ParseInt(s, 0, 0)
		if err != nil {
			return nil, err
		}
		result = append(result, int(converted))
	}

	return result, nil
}
