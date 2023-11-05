package solutions

import (
	"log"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type numRange struct{ start, end int }

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
	allPoints := []int{ranges[0].start, ranges[0].end, ranges[1].start, ranges[1].end}
	slices.Sort(allPoints)
	middleRange := numRange{allPoints[1], allPoints[2]}
	return reflect.DeepEqual(middleRange, ranges[0]) ||
		reflect.DeepEqual(middleRange, ranges[1])
}
