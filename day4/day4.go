package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

/*
*
For all the pairs, how many have a range that fully contains the other?
2-8,3-7 -> 2, 3, 4, 5, 6, 7, 8 & 3, 4, 5, 6, 7
This is an example where the first range fully contains the second.
Only concerned with fully contained, not partially contained.

char 0 & 2 -> (a, b)
char 4 & 6 -> (c, d)
*/
func part1(lines []string) {
	totalContained := 0
	for _, line := range lines {
		// to get ranges, split at ','
		ranges := strings.Split(line, ",")
		// to get start, end, split at '-'
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		range1Start, err := strconv.Atoi(range1[0])
		range1End, err := strconv.Atoi(range1[1])
		range2Start, err := strconv.Atoi(range2[0])
		range2End, err := strconv.Atoi(range2[1])

		if err != nil {
			panic(err)
		}

		if range1Start <= range2Start && range1End >= range2End ||
			range2Start <= range1Start && range2End >= range1End {
			totalContained++
		}
	}
	fmt.Println("Total contained: ", totalContained)
}

/*
*
Find the number of pairs that overlap at all
*/
func part2(lines []string) {
	totalOverlap := 0
	for _, line := range lines {
		// to get ranges, split at ','
		ranges := strings.Split(line, ",")
		// to get start, end, split at '-'
		range1 := strings.Split(ranges[0], "-")
		range2 := strings.Split(ranges[1], "-")

		range1Start, err := strconv.Atoi(range1[0])
		range1End, err := strconv.Atoi(range1[1])
		range2Start, err := strconv.Atoi(range2[0])
		range2End, err := strconv.Atoi(range2[1])

		if err != nil {
			panic(err)
		}

		// if any of the numbers overlap, increment totalOverlap
		if range1Start <= range2Start && range1End >= range2Start ||
			range1Start <= range2End && range1End >= range2End ||
			range2Start <= range1Start && range2End >= range1Start ||
			range2Start <= range1End && range2End >= range1End {
			totalOverlap++
		}
	}
	fmt.Println("Total overlap: ", totalOverlap)
}

func main() {
	lines := utils.Reader("day4.txt")
	fmt.Println("Day 4, Part 1")
	part1(lines)
	fmt.Println()
	fmt.Println("Day 4, Part 2")
	part2(lines)
}
