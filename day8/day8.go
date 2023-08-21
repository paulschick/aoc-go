package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

const (
	Left   = "left"
	Right  = "right"
	Top    = "top"
	Bottom = "bottom"
)

func part1(lines []string) int {
	grid := parseInput(lines)
	visibleCoords := make(map[[2]int]string)

	// row iteration, left/right
	for row := 1; row < len(grid)-1; row++ {
		leftHigh, rightHigh := -1, -1
		leftCoord, rightCoord := [2]int{}, [2]int{}

		for col := 0; col < len(grid[row])-1; col++ {
			height := grid[row][col]
			if height > leftHigh {
				leftHigh = height
				leftCoord = [2]int{row, col}
			}
			rCol := len(grid[row]) - 1 - col
			rHeight := grid[row][rCol]
			if rHeight > rightHigh {
				rightHigh = rHeight
				rightCoord = [2]int{row, rCol}
			}
			visibleCoords[leftCoord] = Left
			visibleCoords[rightCoord] = Right
		}
	}

	// col iteration, top/bottom
	for col := 1; col < len(grid[0])-1; col++ {
		topHigh, bottomHigh := -1, -1
		topCoord, bottomCoord := [2]int{}, [2]int{}

		for row := 0; row < len(grid)-1; row++ {
			height := grid[row][col]
			if height > topHigh {
				topHigh = height
				topCoord = [2]int{row, col}
			}
			bRow := len(grid) - 1 - row
			bHeight := grid[bRow][col]
			if bHeight > bottomHigh {
				bottomHigh = bHeight
				bottomCoord = [2]int{bRow, col}
			}
			visibleCoords[topCoord] = Top
			visibleCoords[bottomCoord] = Bottom
		}
	}

	count := len(visibleCoords)

	// 4 corners
	return count + 4
}

func part2(lines []string) int {
	grid := parseInput(lines)
	highestCount := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			// Need to look at each direction and count the number of numbers
			// that are lower than the current number and stop when we
			// encounter a number that is higher than the current number or the same,
			// count this number and then move on

			currentValue := grid[row][col]
			countLeft, countRight, countTop, countBottom := 0, 0, 0, 0

			// left
			for i := col - 1; i >= 0; i-- {
				if grid[row][i] >= currentValue {
					countLeft++
					break
				}
				countLeft++
			}
			// right
			for i := col + 1; i < len(grid[row]); i++ {
				if grid[row][i] >= currentValue {
					countRight++
					break
				}
				countRight++
			}
			// top
			for i := row - 1; i >= 0; i-- {
				if grid[i][col] >= currentValue {
					countTop++
					break
				}
				countTop++
			}
			// bottom
			for i := row + 1; i < len(grid); i++ {
				if grid[i][col] >= currentValue {
					countBottom++
					break
				}
				countBottom++
			}
			count := countLeft * countRight * countTop * countBottom
			if count > highestCount {
				highestCount = count
			}
		}
	}
	return highestCount
}

func main() {
	fmt.Println("Day 8, Part 1")
	lines := utils.Reader("day8.txt")
	part1Result := part1(lines)
	fmt.Println(part1Result)
	part2Result := part2(lines)
	fmt.Println(part2Result)
}

func parseInput(lines []string) [][]int {
	matrix := make([][]int, len(lines))
	for i, line := range lines {
		matrix[i] = make([]int, len(line))
		for j, character := range line {
			value, err := strconv.Atoi(string(character))
			if err != nil {
				panic(err)
			}
			matrix[i][j] = value
		}
	}
	return matrix
}
