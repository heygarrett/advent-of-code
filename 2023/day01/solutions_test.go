package solutions_test

import (
	solutions "2023/day01"
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"
)

var testCases = []struct {
	input  string
	output int
}{
	{input: "1abc2", output: 12},
	{input: "pqr3stu8vwx", output: 38},
	{input: "a1b2c3d4e5f", output: 15},
	{input: "treb7uchet", output: 77},
}

func TestPart1(t *testing.T) {
	t.Run("example.txt", func(t *testing.T) {
		rawInput, err := fs.ReadFile(os.DirFS("."), "example.txt")
		if err != nil {
			log.Fatalln(err)
		}
		input := string(rawInput)

		got := solutions.Part1(input)
		want := 142

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})

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
