package solutions

import (
	"log"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type numRange struct{ start, end int }

func (nr numRange) getPoints() []int {
	return []int{nr.start, nr.end}
}

func (nr numRange) isEqualTo(newRange numRange) bool {
	return reflect.DeepEqual(nr, newRange)
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	overlaps := 0
	for _, l := range lines {
		firstRange, secondRange := getRangesFromLine(l)
		rangeOverlap := getCompleteOverlap(firstRange, secondRange)
		if rangeOverlap {
			overlaps += 1
		}
	}

	return overlaps
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	overlaps := 0
	for _, l := range lines {
		firstRange, secondRange := getRangesFromLine(l)
		rangeOverlap := getOverlap(firstRange, secondRange)
		if rangeOverlap {
			overlaps += 1
		}
	}

	return overlaps
}

func getRangesFromLine(line string) (numRange, numRange) {
	strs := strings.Split(line, ",")

	ranges := make([]numRange, 0)
	for _, strRange := range strs {
		nums := strings.Split(strRange, "-")
		start, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalln(err)
		}
		end, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalln(err)
		}
		ranges = append(ranges, numRange{start, end})
	}

	return ranges[0], ranges[1]
}

func getCompleteOverlap(firstRange, secondRange numRange) bool {
	allPoints := []int{}
	allPoints = append(allPoints, firstRange.getPoints()...)
	allPoints = append(allPoints, secondRange.getPoints()...)
	slices.Sort(allPoints)
	middleRange := numRange{allPoints[1], allPoints[2]}
	return middleRange.isEqualTo(firstRange) || middleRange.isEqualTo(secondRange)
}

func getOverlap(firstRange, secondRange numRange) bool {
	if getCompleteOverlap(firstRange, secondRange) {
		return true
	}

	rangeEndsDiff := firstRange.end - secondRange.end
	if rangeEndsDiff == 0 {
		return true
	} else if rangeEndsDiff > 0 {
		if firstRange.start <= secondRange.end {
			return true
		}
	} else {
		if secondRange.start <= firstRange.end {
			return true
		}
	}

	return false
}
