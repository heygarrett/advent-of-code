package solutions_test

import (
	solutions "2023/day03"
	"io/fs"
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	rawInput, err := fs.ReadFile(os.DirFS("."), "example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	input := string(rawInput)

	got := solutions.Part1(input)
	want := 4361

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
