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
		total += parseLine(scanner.Text())
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

func parseLine(line string) int {
	digits := filterOutLetters(line)
	numberStr := getFirstAndLastDigit(digits)
	result, err := strconv.ParseInt(numberStr, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}

	return int(result)
}

func filterOutLetters(input string) string {
	re := regexp.MustCompile(`\d`)
	match := re.FindAllString(input, -1)
	digits := strings.Join(match, "")

	return digits
}

func getFirstAndLastDigit(digits string) string {
	firstDigit := digits[:1]
	lastDigit := digits[len(digits)-1:]
	joinedDigits := strings.Join([]string{firstDigit, lastDigit}, "")

	return joinedDigits
}
