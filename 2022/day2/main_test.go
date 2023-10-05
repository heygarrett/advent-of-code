package main_test

import (
	reader "2022"
	"2022/day2"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

var cases = []struct {
	Input       string
	Part1Output int
	Part2Output int
}{
	{
		Input:       getExampleFileInput(),
		Part1Output: 15,
		Part2Output: 12,
	},
	{
		Input: strings.Join([]string{
			"A Y",
			"C Z",
			"B X",
			"A X",
			"B Y",
			"C Z",
			"A Y",
			"B X",
		}, "\n"),
		Part1Output: 39,
		Part2Output: 35,
	},
}

func getExampleFileInput() string {
	input, err := reader.ReadFileFromFS(os.DirFS("."), "example.txt")
	if err != nil {
		log.Fatal(err)
	}

	return input
}

func TestGetTotalScorePart1(t *testing.T) {
	for index, c := range cases {
		t.Run(fmt.Sprintf("case %d", index+1), func(t *testing.T) {
			got := main.GetTotalScorePart1(c.Input)
			want := c.Part1Output

			if got != want {
				t.Errorf("got %d, wanted %d", got, want)
			}
		})
	}
}

func TestGetTotalScorePart2(t *testing.T) {
	for index, c := range cases {
		t.Run(fmt.Sprintf("case %d", index+1), func(t *testing.T) {
			got := main.GetTotalScorePart2(c.Input)
			want := c.Part2Output

			if got != want {
				t.Errorf("got %d, wanted %d", got, want)
			}
		})
	}
}
