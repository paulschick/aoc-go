package utils

import (
	"bufio"
	"fmt"
	"os"
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
