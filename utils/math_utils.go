package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MinInt(nums ...int) int {
	minNum := nums[0]
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
	}
	return minNum
}

// handle file.Close potential error
// allow use of defer
func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Reader(fileName string) []string {
	f, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer closeFile(f)

	scanner := bufio.NewScanner(f)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return lines
}

// Abs returns the absolute value of a number
func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// ToInt converts a string to an int
func ToInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return num
}

func SortInts(ints []int) {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i] < ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
}
