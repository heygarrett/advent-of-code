package main

import (
	reader "2022"
	solutions "2022/day04"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := reader.ReadFileFromFS(os.DirFS("day04"), "input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	part1 := solutions.Part1(input)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := solutions.Part2(input)
	fmt.Printf("Part 2: %d\n", part2)
}
