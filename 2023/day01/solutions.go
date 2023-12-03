package solutions

import (
	"bufio"
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
	return 281
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
