package reader

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

func ReadFileFromFS(fileSystem fs.FS, fileName string) (string, error) {
	file, err := fileSystem.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return readFile(file), nil
}

func readFile(file io.Reader) string {
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "\n")
}
