package main_test

import (
	reader "2022"
	"2022/day01"
	"os"
	"strings"
	"testing"
)

func TestGetMaxCalories(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		calories := strings.Join([]string{
			"1000",
			"2000",
			"",
			"3000",
			"4000",
			"",
			"5000",
			"6000",
			"",
			"7000",
			"8000",
		}, "\n")
		got, err := main.GetMaxCalories(calories)
		if err != nil {
			t.Error(err)
		}

		want := 15000

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

	t.Run("Test 2", func(t *testing.T) {
		calories := strings.Join([]string{"10000", "2000", "", "3000", "4000"}, "\n")
		got, err := main.GetMaxCalories(calories)
		if err != nil {
			t.Error(err)
		}

		want := 12000

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
	t.Run("Example input", func(t *testing.T) {
		input, err := reader.ReadFileFromFS(os.DirFS("."), "example.txt")
		if err != nil {
			t.Error(err)
		}

		got, err := main.GetMaxCalories(input)
		if err != nil {
			t.Error(err)
		}

		want := 24000

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}

func TestGetTotalCaloriesTopThree(t *testing.T) {
	input, err := reader.ReadFileFromFS(os.DirFS("."), "example.txt")
	if err != nil {
		t.Error(err)
	}

	got, err := main.GetTotalCaloriesTopThree(input)
	if err != nil {
		t.Error(err)
	}
	want := 45000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
