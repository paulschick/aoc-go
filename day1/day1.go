package main

import (
	"aoc/input_data_reader"
	"fmt"
	"strconv"
)

func part1(lines []string) {
	maxValue := 0
	groupValue := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || i == len(lines)-1 {
			if groupValue > maxValue {
				maxValue = groupValue
			}
			groupValue = 0
			continue
		}

		value, _ := strconv.Atoi(lines[i])
		groupValue += value
	}

	fmt.Println(maxValue)
}

func main() {
	lines := input_data_reader.Reader("day1.txt")
	fmt.Println("Day 1, part 1")
	part1(lines)
}
