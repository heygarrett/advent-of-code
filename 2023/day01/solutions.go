package solutions

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string) int {
	var total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		digits := findDigits(scanner.Text())
		result, err := strconv.ParseInt(digits[:1]+digits[len(digits)-1:], 0, 0)
		if err != nil {
			log.Fatalln(err)
		}
		total += int(result)
	}

	return total
}

func Part2(input string) int {
	var total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		for key, value := range replacements {
			line = strings.ReplaceAll(line, key, fmt.Sprintf("%s%d", key, value))
		}
		total += parseLine(line)
	}

	return total
}

var replacements = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findDigits(input string) string {
	re := regexp.MustCompile(`\d`)
	match := re.FindAllString(input, -1)
	digits := strings.Join(match, "")

	return digits
}

