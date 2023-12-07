package solutions

import (
	"bufio"
	"log"
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

func Part1(input string) int {
	var total int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		newGame := parseLine(scanner.Text())
		if newGame.isPossible(cubeSet{red: 13, green: 14, blue: 15}) {
			total += newGame.id
		}
	}

	return total
}

func parseLine(line string) game {
	splitOnColon := strings.Split(line, ": ")
	gameIdStr, gameSetsStr := splitOnColon[0], splitOnColon[1]
	gameSet := strings.Split(gameSetsStr, "; ")

	idStr := strings.Split(gameIdStr, " ")[1]
	id, err := strconv.ParseInt(idStr, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	cubeSets := make([]cubeSet, 0)
	for _, set := range gameSet {
		newCubeSet := cubeSet{}
		for _, css := range strings.Split(set, ", ") {
			numberAndColor := strings.Split(css, " ")
			color := numberAndColor[1]
			number, err := strconv.ParseInt(numberAndColor[0], 0, 0)
			if err != nil {
				log.Fatalln(err)
			}
			switch color {
			case "red":
				newCubeSet.red = int(number)
			case "green":
				newCubeSet.green = int(number)
			case "blue":
				newCubeSet.blue = int(number)
			}
		}
		cubeSets = append(cubeSets, newCubeSet)
	}

	return game{id: int(id), sets: cubeSets}
}
