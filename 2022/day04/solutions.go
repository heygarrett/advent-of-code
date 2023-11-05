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
		ranges := getRangesFromLine(l)
		rangeOverlap := getOverlap(ranges)
		if rangeOverlap {
			overlaps += 1
		}
	}

	return overlaps
}

func getRangesFromLine(line string) []numRange {
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

	return ranges
}

func getOverlap(ranges []numRange) bool {
	allPoints := []int{}
	allPoints = append(allPoints, ranges[0].getPoints()...)
	allPoints = append(allPoints, ranges[1].getPoints()...)
	slices.Sort(allPoints)
	middleRange := numRange{allPoints[1], allPoints[2]}
	return middleRange.isEqualTo(ranges[0]) || middleRange.isEqualTo(ranges[1])
}
