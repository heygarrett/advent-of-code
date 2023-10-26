package main

import (
	reader "2022"
	solutions "2022/day03"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := reader.ReadFileFromFS(os.DirFS("day03"), "input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1 := solutions.Part1(input)
	fmt.Printf("Part 1: %d\n", part1)
}
