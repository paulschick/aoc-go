package main

import (
	"aoc/input_data_reader"
	"fmt"
	"strings"
	"unicode"
)

/**
Each of the lines represents two sets of characters.
Uppercase and lowercase letters are different

Each set has the same number of characters, so the first half is the first set
and the second half is the second set.

part 1

One character will appear in both sets for each line. You need to get the priority number for
the repeated character in each line, and then get the sum for each line.

a-z = 1-26
A-Z = 27-52

Capital letters - subtract 38 from the ascii value
Lowercase letters - subtract 96 from the ascii value
*/

func part1(lines []string) {
	totalSum := 0

	for _, line := range lines {
		lineLen := len(line)
		partLength := lineLen / 2
		linePart1 := line[:partLength]
		linePart2 := line[partLength:]

		// only do this while there's no match
		for _, char := range linePart1 {
			// is this char contained in linePart2?
			if strings.Contains(linePart2, string(char)) {
				if unicode.IsUpper(char) {
					totalSum += int(char) - 38
				} else {
					totalSum += int(char) - 96
				}
				break
			}
		}
	}

	fmt.Println("total sum: ", totalSum)
}

/*
*
Now do in groups of three lines instead of within each line
*/
func part2(lines []string) {
	totalSum := 0

	for i := 0; i < len(lines); i += 3 {
		groupSum := 0
		lineGroup := lines[i : i+3]
		line1 := lineGroup[0]
		for _, char := range line1 {
			if strings.Contains(lineGroup[1], string(char)) && strings.Contains(lineGroup[2], string(char)) {
				if unicode.IsUpper(char) {
					groupSum += int(char) - 38
				} else {
					groupSum += int(char) - 96
				}
				break
			}
		}
		totalSum += groupSum
	}
	fmt.Println("total sum: ", totalSum)
}

func main() {
	lines := input_data_reader.Reader("day3.txt")
	fmt.Println("day 3, part 1")
	part1(lines)
	fmt.Println("day 3, part 2")
	part2(lines)
}
