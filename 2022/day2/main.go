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

	totalScore := GetTotalScorePart1(input)
	fmt.Printf("Total score, part 1: %d\n", totalScore)
	totalScore = GetTotalScorePart2(input)
	fmt.Printf("Total score, part 2: %d\n", totalScore)
}

func GetTotalScorePart1(input string) int {
	parsedInput := parseInput(input)
	totalScore := 0
	for _, round := range parsedInput {
		opponentShape := shapes[round[0]]
		myShape := shapes[round[1]]
		outcome := getOutcomeFromShapes(myShape, opponentShape)
		totalScore += int(myShape) + int(outcome)
	}

	return totalScore
}

func GetTotalScorePart2(input string) int {
	parsedInput := parseInput(input)
	totalScore := 0
	for _, round := range parsedInput {
		opponentShape := shapes[round[0]]
		outcome := outcomes[round[1]]
		myShape := getShapeFromOutcome(opponentShape, outcome)
		totalScore += int(myShape) + int(outcome)
	}

	return totalScore
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	letterPairs := make([][]string, 0)
	for _, line := range lines {
		letterPairs = append(letterPairs, strings.Split(line, " "))
	}

	return letterPairs
}

type Shape int

const (
	rock     Shape = 1
	paper    Shape = 2
	scissors Shape = 3
)

var shapes = map[string]Shape{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var losesAgainst = map[Shape]Shape{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

var winsAgainst = map[Shape]Shape{
	rock:     paper,
	paper:    scissors,
	scissors: rock,
}

type Outcome int

const (
	loss Outcome = 0
	draw Outcome = 3
	win  Outcome = 6
)

var outcomes = map[string]Outcome{
	"X": loss,
	"Y": draw,
	"Z": win,
}

func getOutcomeFromShapes(myShape, opponentShape Shape) Outcome {
	if myShape == losesAgainst[opponentShape] {
		return loss
	} else if myShape == opponentShape {
		return draw
	} else {
		return win
	}
}

func getShapeFromOutcome(opponentShape Shape, outcome Outcome) Shape {
	if outcome == draw {
		return opponentShape
	} else if outcome == loss {
		return losesAgainst[opponentShape]
	} else {
		return winsAgainst[opponentShape]
	}
}
