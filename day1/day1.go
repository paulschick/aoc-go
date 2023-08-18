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

func part2(lines []string) {
	top1 := 0
	top2 := 0
	top3 := 0
	sumOf3 := 0
	groupValue := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || i == len(lines)-1 {
			if groupValue > top3 || groupValue > top2 || groupValue > top1 {
				if groupValue > top1 {
					top3 = top2
					top2 = top1
					top1 = groupValue
				} else if groupValue > top2 {
					top3 = top2
					top2 = groupValue
				} else if groupValue > top3 {
					top3 = groupValue
				}
				sumOf3 = top1 + top2 + top3
			}
			groupValue = 0
			continue
		}

		value, _ := strconv.Atoi(lines[i])
		groupValue += value
	}
	fmt.Println(sumOf3)
}

func main() {
	lines := input_data_reader.Reader("day1.txt")
	fmt.Println("Day 1, part 1")
	part1(lines)
	fmt.Println("Day 1, part 2")
	part2(lines)
}
