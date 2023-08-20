package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Dir struct {
	name      string
	parentDir *Dir
	Children  map[string]*Dir
	Files     map[string]int
	totalSize int
}

func part1(lines []string) int {
	root := parseInput(lines)
	return sumUnder100k(root)
}

/*
* Max Space: 70000000
* Minimum unused for update: 30000000
* Delete a directory to allow enough space. Out of all directories which are large enough to provide
* this space, choose the smallest one.
* Return the total size of the deleted directory.
 */
func part2(lines []string) int {
	root := parseInput(lines)
	totalAvailable := 70000000
	minUnused := 30000000
	minimumSize := minUnused - (totalAvailable - root.totalSize)
	return findSmallestToDelete(root, minimumSize)
}

// Reference:
// https://github.com/alexchao26/advent-of-code-go/blob/main/2022/day07/main.go
func parseInput(lines []string) *Dir {
	root := &Dir{
		name:     "root",
		Children: map[string]*Dir{},
	}
	currentDir := root

	i := 0
	for i < len(lines) {
		switch cmd := lines[i]; cmd[0:1] {
		case "$":
			if cmd == "$ ls" {
				i++
			} else {
				changeDir := strings.Split(cmd, "cd ")[1]
				if changeDir == ".." {
					if currentDir.parentDir == nil {
						i++
						continue
					}
					currentDir = currentDir.parentDir
				} else {
					// if changeDir doesn't exist, create it and set as current
					if _, found := currentDir.Children[changeDir]; !found {
						currentDir.Children[changeDir] = &Dir{
							name:      changeDir,
							parentDir: currentDir,
							Children:  map[string]*Dir{},
							Files:     map[string]int{},
						}
					}

					// set current directory to changeDir
					currentDir = currentDir.Children[changeDir]
				}
				i++
			}
		// default case is listing contents of a directory
		default:
			// if it's a dir, we need to see if it's listed in the children.
			// if not, create it.
			if strings.HasPrefix(cmd, "dir") {
				name := cmd[4:]
				if _, found := currentDir.Children[name]; !found {
					currentDir.Children[name] = &Dir{
						name:      name,
						parentDir: currentDir,
						Children:  map[string]*Dir{},
						Files:     map[string]int{},
					}
				}
			} else {
				parts := strings.Split(cmd, " ")
				name := parts[1]
				size := parts[0]
				sizeInt, err := strconv.Atoi(size)
				if err != nil {
					panic("Error converting string to int")
				}
				currentDir.Files[name] = sizeInt
			}
			i++
		}
	}
	setFileSizes(root)
	return root
}

func setFileSizes(currentDir *Dir) int {
	totalSize := 0
	for _, file := range currentDir.Files {
		totalSize += file
	}
	for _, child := range currentDir.Children {
		totalSize += setFileSizes(child)
	}
	currentDir.totalSize = totalSize
	return totalSize
}

func sumUnder100k(dir *Dir) int {
	sum := 0
	if dir.totalSize < 100000 {
		sum += dir.totalSize
	}
	for _, child := range dir.Children {
		sum += sumUnder100k(child)
	}
	return sum
}

func findSmallestToDelete(dir *Dir, dirMinSize int) int {
	smallest := math.MaxInt64

	if dir.totalSize >= dirMinSize {
		smallest = utils.MinInt(smallest, dir.totalSize)
	}

	for _, child := range dir.Children {
		smallest = utils.MinInt(smallest, findSmallestToDelete(child, dirMinSize))
	}
	return smallest
}

func main() {
	lines := utils.Reader("day7.txt")
	fmt.Println("Day 1, Part 1")
	under100kSum := part1(lines)
	fmt.Println()
	fmt.Println("Sum of files under 100k: ", under100kSum)
	fmt.Println()
	fmt.Println("Day 1, Part 2")
	smallestToDelete := part2(lines)
	fmt.Println("Smallest to delete: ", smallestToDelete)
}
