package solutions

import (
	"bufio"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type matchers []func(string) string

func Part1(input string) int {
	var total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		firstNumber := getFirstNumber(line, matchers{matchDigit})
		lastNumber := getLastNumber(line, matchers{matchDigit})
		total += convertToInt(firstNumber, lastNumber)
	}

	return total
}

func Part2(input string) int {
	var total int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		firstNumber := getFirstNumber(line, matchers{matchDigit, matchWord})
		lastNumber := getLastNumber(line, matchers{matchDigit, matchWord})
		total += convertToInt(firstNumber, lastNumber)
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

func convertToInt(firstDigit, lastDigit string) int {
	result, err := strconv.ParseInt(firstDigit+lastDigit, 0, 0)
	if err != nil {
		log.Fatalln(err)
	}
	return int(result)
}

func getFirstNumber(input string, callbacks matchers) string {
	for end := 1; end <= len(input); end++ {
		strSlice := input[:end]
		for _, f := range callbacks {
			if match := f(strSlice); match != "" {
				return match
			}
		}
	}

	return ""
}

func getLastNumber(input string, callbacks matchers) string {
	for start := len(input) - 1; start >= 0; start-- {
		strSlice := input[start:]
		for _, f := range callbacks {
			if match := f(strSlice); match != "" {
				return match
			}
		}
	}

	return ""
}

func matchDigit(input string) string {
	re := regexp.MustCompile(`\d`)
	match := re.FindString(input)
	return match
}

func matchWord(input string) string {
	for k, v := range replacements {
		if strings.Contains(input, k) {
			return v
		}
	}

	return ""
}
