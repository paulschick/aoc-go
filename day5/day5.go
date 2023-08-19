package main

import (
	"aoc/input_data_reader"
	"day_5/lib"
	"fmt"
)

func part1(lines []string) {
	//columns := constructStacks(lines)

	//for _, column := range columns {
	// print the top of the stack
	//value, _ := column.Pop()
	//fmt.Println(value)
	//}

	instructions := lib.ParseInstructions(lines)
	for _, instruction := range instructions {
		fmt.Println(instruction)
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
