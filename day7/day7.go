package main

import (
	"aoc/input_data_reader"
	"day_7/node_tree"
	"fmt"
)

func part1(lines []string) {
	node_tree.ProcessTree(lines)
}

func part2(lines []string) {
	fmt.Println("not implemented")
}

func main() {
	lines := input_data_reader.Reader("day7.txt")
	fmt.Println("Day 7, Part 1")
	part1(lines)
	fmt.Println("Day 7, Part 2")
	part2(lines)
}
