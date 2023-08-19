package main

import (
	"aoc/input_data_reader"
	"day_5/lib"
	"day_5/types"
	"fmt"
)

func part1(lines []string) {
	columns := lib.ConstructStacks(lines)
	instructions := lib.ParseInstructions(lines)
	fmt.Println("Columns: ", columns)

	for i := 0; i < len(instructions); i++ {
		// Need to use a pointer to actually update the column
		fromColumn := &columns[instructions[i].FromIndex]
		toColumn := &columns[instructions[i].ToIndex]

		for a := 1; a <= instructions[i].Number; a++ {
			fromValue, isValid := fromColumn.Pop()

			if isValid {
				toColumn.Push(fromValue)
			}
		}
	}
	fmt.Println("Columns: ", columns)
}

func part2(lines []string) {
	columns := lib.ConstructStacks(lines)
	instructions := lib.ParseInstructions(lines)
	fmt.Println("Columns: ", columns)
	for i := 0; i < len(instructions); i++ {
		fmt.Println("instruction: ", instructions[i])

		fromColumn := &columns[instructions[i].FromIndex]
		toColumn := &columns[instructions[i].ToIndex]

		var fromPartialStack types.Stack

		for a := 1; a <= instructions[i].Number; a++ {
			fromValue, isValid := fromColumn.Pop()

			if isValid {
				fromPartialStack.PushBottom(fromValue)
			}
		}

		// concatentate the partial stack to the toColumn
		for j := 0; j < len(fromPartialStack); j++ {
			toColumn.Push(fromPartialStack[j])
		}
	}
	for _, column := range columns {
		fmt.Println(column)
	}
}

func main() {
	lines := input_data_reader.Reader("day5.txt")
	fmt.Println("Day 5, Part 1")
	part1(lines)
	fmt.Println("Day 5, Part 2")
	part2(lines)
}
