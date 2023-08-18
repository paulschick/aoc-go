package main

import (
	"aoc/input_data_reader"
	"fmt"
)

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (string, bool) {
	l := len(*s)
	if l == 0 {
		return "", false
	}
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res, true
}

func part1(lines []string) {
	// top to bottom of the stacks
	// going by row
	for _, line := range lines {
		if line == "" || line[1] == '1' {
			break
		}
		var lineValues []string
		for i := 1; i < len(line); i += 4 {
			// last line, pad with spaces
			if i+4 > len(line) {
				lineValues = append(lineValues, string(line[i]))
				remainingColumns := 9 - len(lineValues)
				for j := 0; j < remainingColumns; j++ {
					lineValues = append(lineValues, " ")
				}
				break
			}
			lineValues = append(lineValues, string(line[i]))
		}
		fmt.Println(lineValues)
	}
}

func part2(lines []string) {
	fmt.Println("not implemented")
}

func main() {
	lines := input_data_reader.Reader("day5.txt")
	fmt.Println("Day 5, Part 1")
	part1(lines)
	fmt.Println("Day 5, Part 2")
	part2(lines)
}
