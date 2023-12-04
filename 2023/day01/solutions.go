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
		firstNumber := getFirstNumber(line)
		lastNumber := getLastNumber(line)
		result, err := strconv.ParseInt(firstNumber+lastNumber, 0, 0)
		if err != nil {
			log.Fatalln(err)
		}
		total += int(result)
	}

	return total
}

var replacements = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getFirstNumber(input string) string {
	for end := 1; end <= len(input); end++ {
		strSlice := input[:end]
		re := regexp.MustCompile(`\d`)
		match := re.FindString(strSlice)
		if match != "" {
			return match
		}

		for k, v := range replacements {
			if strings.Contains(strSlice, k) {
				return v
			}
		}
	}

	return ""
}

func getLastNumber(input string) string {
	for start := len(input) - 1; start >= 0; start-- {
		strSlice := input[start:]
		re := regexp.MustCompile(`\d`)
		match := re.FindString(strSlice)
		if match != "" {
			return match
		}

		for k, v := range replacements {
			if strings.Contains(strSlice, k) {
				return v
			}
		}
	}

	return ""
}

func findDigits(input string) string {
	re := regexp.MustCompile(`\d`)
	match := re.FindAllString(input, -1)
	digits := strings.Join(match, "")

	return digits
}
