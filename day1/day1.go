package main

import (
	"aoc/input_data_reader"
	"fmt"
)

func main() {
	lines := input_data_reader.Reader("day1.txt")

	for _, value := range lines {
		fmt.Println(value)
	}
}
