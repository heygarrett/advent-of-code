package solutions_test

import (
	solutions "2023/day02"
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"
)

func getExampleInput() string {
	rawInput, err := fs.ReadFile(os.DirFS("."), "example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(rawInput)

	return input
}

func TestPart1(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		got := solutions.Part1(getExampleInput())
		want := 8

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
	testCases := []struct {
		input  string
		output int
	}{
		{
			input:  "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			output: 1,
		},
		{
			input:  "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			output: 2,
		},
		{
			input:  "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			output: 0,
		},
		{
			input:  "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			output: 0,
		},
		{
			input:  "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			output: 5,
		},
	}

	for index, tc := range testCases {
		t.Run(fmt.Sprintf("Test case: %d", index+1), func(t *testing.T) {
			got := solutions.Part1(tc.input)
			want := tc.output

			if got != want {
				t.Errorf("got %d, wanted %d", got, want)
			}
		})
	}
}
