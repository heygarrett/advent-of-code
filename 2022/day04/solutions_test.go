package solutions_test

import (
	reader "2022"
	solutions "2022/day04"
	"fmt"
	"log"
	"os"
	"testing"
)

func getExampleInput() string {
	input, err := reader.ReadFileFromFS(os.DirFS("."), "example.txt")
	if err != nil {
		log.Fatal(err)
	}

	return input
}

type Output struct {
	part1, part2 int
}

var cases = []struct {
	input  string
	output Output
}{
	{input: "2-4,6-8", output: Output{part1: 0, part2: 0}},
	{input: "2-3,4-5", output: Output{part1: 0, part2: 0}},
	{input: "5-7,7-9", output: Output{part1: 0, part2: 1}},
	{input: "2-8,3-7", output: Output{part1: 1, part2: 1}},
	{input: "6-6,4-6", output: Output{part1: 1, part2: 1}},
	{input: "2-6,4-8", output: Output{part1: 0, part2: 1}},
}

func TestPart1(t *testing.T) {
	for index, test := range cases {
		t.Run(fmt.Sprintf("Test case %d", index+1), func(t *testing.T) {
			got := solutions.Part1(test.input)
			want := test.output.part1

			if got != want {
				t.Errorf("got %d, wanted %d", got, want)
			}
		})
	}

	t.Run("example.txt", func(t *testing.T) {
		input := getExampleInput()

		got := solutions.Part1(input)
		want := 2

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	for index, test := range cases {
		t.Run(fmt.Sprintf("Test case %d", index+1), func(t *testing.T) {
			got := solutions.Part2(test.input)
			want := test.output.part2

			if got != want {
				t.Errorf("got %d, wanted %d", got, want)
			}
		})
	}

	t.Run("example.txt", func(t *testing.T) {
		input := getExampleInput()

		got := solutions.Part2(input)
		want := 4

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}
