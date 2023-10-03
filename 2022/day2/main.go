package main

import (
	reader "2022"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := reader.ReadFileFromFS(os.DirFS("day2"), "input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	totalScore := GetTotalScore(input)
	fmt.Printf("Total score: %d\n", totalScore)
}

func GetTotalScore(input string) int {
	parsedInput := parseInput(input)
	totalScore := 0
	for _, round := range parsedInput {
		totalScore += getRoundScore(round)
	}

	return totalScore
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

type Shape int

const (
	rock Shape = iota
	paper
	scissors
)

func getRoundScore(round string) int {
	shapes := map[string]Shape{
		"A": rock,
		"B": paper,
		"C": scissors,
		"X": rock,
		"Y": paper,
		"Z": scissors,
	}

	bothShapes := strings.Split(round, " ")
	opponentShape := shapes[bothShapes[0]]
	myShape := shapes[bothShapes[1]]

	myShapeScore := getShapeScore(myShape)
	outcomeScore := getOutcomeScore(myShape, opponentShape)

	return myShapeScore + outcomeScore
}

func getShapeScore(shape Shape) int {
	shapeValues := map[Shape]int{
		rock:     1,
		paper:    2,
		scissors: 3,
	}

	return shapeValues[shape]
}

func getOutcomeScore(myShape, opponentShape Shape) int {
	losesAgainst := map[Shape]Shape{
		rock:     scissors,
		paper:    rock,
		scissors: paper,
	}

	if myShape == opponentShape {
		return 3
	} else if myShape == losesAgainst[opponentShape] {
		return 0
	} else {
		return 6
	}
}
