package solutions

import (
	"bufio"
	"log"
	"slices"
	"strconv"
	"strings"
)

type cubeSet struct{ red, green, blue int }

type game struct {
	id   int
	sets []cubeSet
}

func (g game) isPossible(upperLimit cubeSet) bool {
	for _, set := range g.sets {
		if set.red > upperLimit.red {
			return false
		}
		if set.green > upperLimit.green {
			return false
		}
		if set.blue > upperLimit.blue {
			return false
		}
	}
	return true
}

func (g game) getPower() int {
	cubeSet := g.getMinimumSet()
	return cubeSet.red * cubeSet.green * cubeSet.blue
}

func (g game) getMinimumSet() cubeSet {
	minimum := cubeSet{}
	for _, set := range g.sets {
		minimum.red = slices.Max([]int{minimum.red, set.red})
		minimum.green = slices.Max([]int{minimum.green, set.green})
		minimum.blue = slices.Max([]int{minimum.blue, set.blue})
	}

	return minimum
}

func Part1(input string) int {
	return computeLines(input, func(g game) int {
		if g.isPossible(cubeSet{red: 12, green: 13, blue: 14}) {
			return g.id
		} else {
			return 0
		}
	})
}

func Part2(input string) int {
	return computeLines(input, func(g game) int {
		return g.getPower()
	})
}

func computeLines(input string, f func(g game) int) int {
	var total int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		newGame := parseLine(scanner.Text())
		total += f(newGame)
	}

	return total
}

func parseLine(line string) game {
	splitOnColon := strings.Split(line, ": ")
	gameIdString, gameSetString := splitOnColon[0], splitOnColon[1]
	cubeSetStrings := strings.Split(gameSetString, "; ")

	idString := strings.Split(gameIdString, " ")[1]
	idInt64, err := strconv.ParseInt(idString, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}
	idInt := int(idInt64)

	cubeSets := make([]cubeSet, 0)
	for _, setString := range cubeSetStrings {
		newCubeSet := cubeSet{}
		for _, cubeCountString := range strings.Split(setString, ", ") {
			colorAndCount := strings.Split(cubeCountString, " ")
			color := colorAndCount[1]
			count64, err := strconv.ParseInt(colorAndCount[0], 0, 0)
			if err != nil {
				log.Fatalln(err)
			}
			count := int(count64)

			switch color {
			case "red":
				newCubeSet.red = count
			case "green":
				newCubeSet.green = count
			case "blue":
				newCubeSet.blue = count
			}
		}
		cubeSets = append(cubeSets, newCubeSet)
	}

	return game{id: idInt, sets: cubeSets}
}
