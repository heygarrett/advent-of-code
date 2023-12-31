package solutions_test

import (
	reader "2022"
	solutions "2022/day03"
	"fmt"
	"log"
	"os"
	"testing"
)

var cases = []struct {
	input  string
	output int
}{
	{input: "vJrwpWtwJgWrhcsFMMfFFhFp", output: 16},
	{input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", output: 38},
	{input: "PmmdzqPrVvPwwTWBwg", output: 42},
	{input: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", output: 22},
	{input: "ttgJtRGJQctTZtZT", output: 20},
	{input: "CrZsJsPPZsGzwwsLwLmpwMDw", output: 19},
}

func TestPart1(t *testing.T) {
	for index, c := range cases {
		t.Run(fmt.Sprintf("Case %d", index), func(t *testing.T) {
			got := solutions.Part1(c.input)
			want := c.output

			if got != want {
				t.Errorf("got %d, wanted %d", got, want)
			}
		})
	}

	t.Run("example.txt", func(t *testing.T) {
		input := getExampleInput()

		got := solutions.Part1(input)
		want := 157

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {
	input := getExampleInput()

	got := solutions.Part2(input)
	want := 70

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func getExampleInput() string {
	input, err := reader.ReadFileFromFS(os.DirFS("."), "example.txt")
	if err != nil {
		log.Fatal(err)
	}

	return input
}
