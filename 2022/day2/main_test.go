package main_test

import (
	reader "2022"
	main "2022/day2"
	"os"
	"strings"
	"testing"
)

func TestGetTotalScorePart1(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		input, err := reader.ReadFileFromFS(os.DirFS("."), "example.txt")
		if err != nil {
			t.Error(err)
		}

		got := main.GetTotalScorePart1(input)
		want := 15

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("larger example", func(t *testing.T) {
		input := strings.Join([]string{
			"A Y",
			"C Z",
			"B X",
			"A X",
			"B Y",
			"C Z",
			"A Y",
			"B X",
		}, "\n")

		got := main.GetTotalScorePart1(input)
		want := 39

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}

