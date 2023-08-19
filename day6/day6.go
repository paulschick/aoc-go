package main

import (
	"aoc/input_data_reader"
	"fmt"
	"slices"
)

func part1(value string) {
	endIndex := 0
	for i := 0; i < len(value)-3; i++ {
		val1 := value[i]
		val2 := value[i+1]
		val3 := value[i+2]
		val4 := value[i+3]
		if val1 != val2 && val1 != val3 && val1 != val4 && val2 != val3 && val2 != val4 && val3 != val4 {
			fmt.Println("unique valueSet: ", string(val1), string(val2), string(val3), string(val4))
			endIndex = i + 3
			break
		}
	}
	fmt.Println("number of values between beginning of string and endIndex: ", endIndex+1)
}

/*
*
Similar to the first part, but now it's 14 unique characters instead of four
*/
func part2(value string) {
	endIndex := 0
	setCopy := make([]string, 14)
	for i := 0; i < len(value)-13; i++ {
		if endIndex == 0 {
			var valueSet []string
			for j := 0; j < 14; j++ {
				valueSet = append(valueSet, string(value[i+j]))
			}
			copy(setCopy, valueSet)
			validSet := true
			for a := 0; a < 14; a++ {
				if len(valueSet) == 0 {
					break
				}
				currentLength := len(valueSet)
				poppedValue := (valueSet)[currentLength-1]
				valueSet = (valueSet)[:currentLength-1]
				// check if valueSet contains poppedValue
				if slices.Contains(valueSet, poppedValue) {
					validSet = false
					break
				}
			}
			if validSet {
				endIndex = i + 13
			}
		} else {
			break
		}
	}
	fmt.Println("Unique valueSet ", setCopy)
	fmt.Println("number of values between beginning of string and endIndex: ", endIndex+1)
}

/*
*
Given a string of characters, find the first location where 4 characters in a row are all different.
Count the number of characters from the beginning of the string to the end of the four
characters that are different.

example:
`bvwbjplbgvbhsrlpgdmjqwftvncz`
In this, the second character to the 5th character are all different: `vwbj`
So the answer is 5, because that's the number of characters from the start to the end of the four characters
*/
func main() {
	fmt.Println("Day 6, Part 1")
	inputValue := input_data_reader.Reader("day6.txt")[0]
	part1(inputValue)
	fmt.Println("Day 6, Part 2")
	part2(inputValue)
}
