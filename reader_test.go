package reader_test

import (
	reader "2022"
	"testing"
	"testing/fstest"
)

func TestReadFile(t *testing.T) {
	const (
		example1 = "hi there hello"
		example2 = "now you're gtting it"
	)
	fs := fstest.MapFS{
		"example1.txt": {Data: []byte(example1)},
		"example2.txt": {Data: []byte(example2)},
	}

	t.Run("example1", func(t *testing.T) {
		input, err := reader.ReadFileFromFS(fs, "example1.txt")
		if err != nil {
			t.Fatal(err)
		}

		if input != example1 {
			t.Errorf("got %s, wanted %s", input, example1)
		}
	})

	t.Run("example2", func(t *testing.T) {
		input, err := reader.ReadFileFromFS(fs, "example2.txt")
		if err != nil {
			t.Fatal(err)
		}

		if input != example2 {
			t.Errorf("got %s, wanted %s", input, example2)
		}
	})
}
