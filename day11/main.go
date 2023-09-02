package main

import (
	"aoc/utils"
	"day_11/d11_p1"
	"day_11/d11_p2"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing part number argument. Usage: go run main.go part1 or go run main.go part2")
		return
	}

	switch os.Args[1] {
	case "part1":
		d11_p1.Part1(getInput())
	case "part2":
		d11_p2.Part2(getInput())
	default:
		fmt.Printf("Unknown command: %s. Expected 'part1' or 'part2'.\n", os.Args[1])
	}

}

func getInput() []string {
	return utils.Reader("day11.txt")
}
