package main_test

import (
	reader "2022"
	main "2022/day2"
	"os"
	"strings"
	"testing"
)

func TestGetTotalScore(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		input, err := reader.ReadFileFromFS(os.DirFS("."), "example.txt")
		if err != nil {
			t.Error(err)
		}

		got := main.GetTotalScore(input)
		want := 15

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("larger example", func(t *testing.T) {
		input := strings.Join([]string{
			"A Y", // 6 + 2
			"C Z", // 3 + 3
			"B X", // 0 + 1
			"A X", // 3 + 1
			"B Y", // 3 + 2
			"C Z", // 3 + 3
			"A Y", // 6 + 2
			"B X", // 0 + 1
		}, "\n")

		got := main.GetTotalScore(input)
		want := 39

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}
