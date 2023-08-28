package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

const (
	UP    = 1
	DOWN  = 2
	LEFT  = 3
	RIGHT = 4
)

type Instruction struct {
	Direction string
	Steps     int
}

type Point struct {
	X int
	Y int
}

type Knot struct {
	Location         Point
	PreviousLocation Point
	Visited          map[Point]bool
	ParentKnot       *Knot // nil if head
}

func (knot *Knot) moveHead(step int) *Knot {
	knot.PreviousLocation = Point{knot.Location.X, knot.Location.Y}
	knot.Location = moveHead(knot.Location, step)
	return knot
}

func (knot *Knot) moveX(count int) *Knot {
	knot.PreviousLocation = Point{knot.Location.X, knot.Location.Y}
	knot.Location = Point{knot.Location.X + count, knot.Location.Y}
	knot.Visited[knot.Location] = true
	return knot
}

func (knot *Knot) moveY(count int) *Knot {
	knot.PreviousLocation = Point{knot.Location.X, knot.Location.Y}
	knot.Location = Point{knot.Location.X, knot.Location.Y + count}
	knot.Visited[knot.Location] = true
	return knot
}

func (knot *Knot) moveXAndY(countX int, countY int) *Knot {
	knot.PreviousLocation = Point{knot.Location.X, knot.Location.Y}
	knot.Location = Point{knot.Location.X + countX, knot.Location.Y + countY}
	knot.Visited[knot.Location] = true
	return knot
}

func (knot *Knot) moveKnot() *Knot {
	if knot.ParentKnot != nil {
		parentCurrent := Point{knot.ParentKnot.Location.X, knot.ParentKnot.Location.Y}
		xDiff := parentCurrent.X - knot.Location.X
		yDiff := parentCurrent.Y - knot.Location.Y

		if xDiff == -2 && yDiff == 0 {
			knot.moveX(-1)
		} else if xDiff == 2 && yDiff == 0 {
			knot.moveX(1)
		} else if xDiff == 0 && yDiff == -2 {
			knot.moveY(-1)
		} else if xDiff == 0 && yDiff == 2 {
			knot.moveY(1)
		} else if (xDiff == 1 && yDiff == 2) ||
			(xDiff == 2 && yDiff == 1) ||
			(xDiff == 2 && yDiff == 2) {
			knot.moveXAndY(1, 1)
		} else if (xDiff == -1 && yDiff == 2) ||
			(xDiff == -2 && yDiff == 1) ||
			(xDiff == -2 && yDiff == 2) {
			knot.moveXAndY(-1, 1)
		} else if (xDiff == 1 && yDiff == -2) ||
			(xDiff == 2 && yDiff == -1) ||
			(xDiff == 2 && yDiff == -2) {
			knot.moveXAndY(1, -1)
		} else if (xDiff == -1 && yDiff == -2) ||
			(xDiff == -2 && yDiff == -1) ||
			(xDiff == -2 && yDiff == -2) {
			knot.moveXAndY(-1, -1)
		}
	}
	return knot
}

func processGrid(numKnots int) int {
	input := getInput()
	instructions := parseInput(input)

	knots := initializeKnotSlice(numKnots)

	instr := 1
	for _, instruction := range instructions {
		stepSeq := getStepsSequence(instruction)
		for _, step := range stepSeq {
			for i := 0; i < numKnots; i++ {
				if i == 0 {
					headKnot := knots[i]
					knots[0] = *headKnot.moveHead(step)
				} else {
					knot := knots[i]
					knots[i] = *knot.moveKnot()
				}
			}
		}
		instr++
	}
	return len(knots[len(knots)-1].Visited)
}

/*
*
Every time the tail needs to move, it's going to move into the position previously occupied by the head.
This is true for linear and diagonal moves.
*/
func part1() {
	fmt.Println("Day 9, Part 1")
	result := processGrid(2)
	fmt.Println(result)
}

func part2() {
	fmt.Println("Day 9, Part 2")
	result := processGrid(10)
	fmt.Println(result)
}

// getStepsSequence returns a slice of ints that represent the steps to take
func getStepsSequence(instruction Instruction) []int {
	if instruction.Direction == "U" {
		return makeStepSlice(UP, instruction.Steps)
	} else if instruction.Direction == "D" {
		return makeStepSlice(DOWN, instruction.Steps)
	} else if instruction.Direction == "L" {
		return makeStepSlice(LEFT, instruction.Steps)
	} else if instruction.Direction == "R" {
		return makeStepSlice(RIGHT, instruction.Steps)
	}
	return []int{}
}

func makeStepSlice(value int, length int) []int {
	stepsSlice := make([]int, length)
	for i := 0; i < length; i++ {
		stepsSlice[i] = value
	}
	return stepsSlice
}

func getInput() []string {
	return utils.Reader("day9.txt")
}

func parseInput(input []string) []Instruction {
	instructions := make([]Instruction, len(input))
	for i, line := range input {
		instructions[i] = parseInstruction(line)
	}
	return instructions
}

func parseInstruction(line string) Instruction {
	direction := strings.Split(line, " ")[0]
	steps := utils.ToInt(strings.Split(line, " ")[1])
	return Instruction{direction, steps}
}

func moveHead(headLocation Point, step int) Point {
	if step == UP {
		headLocation.Y++
	} else if step == DOWN {
		headLocation.Y--
	} else if step == LEFT {
		headLocation.X--
	} else if step == RIGHT {
		headLocation.X++
	}
	return headLocation
}

func initializeKnotSlice(numKnots int) []Knot {
	// Initialize knots slice
	knots := make([]Knot, numKnots)
	headKnot := Knot{Point{0, 0}, Point{0, 0}, make(map[Point]bool), nil}
	knots[0] = headKnot
	// Create a knot for each number of knots
	// track the position of each knot and its parent in a slice
	for i := 1; i < numKnots; i++ {
		var visited map[Point]bool
		visited = make(map[Point]bool)
		visited[Point{0, 0}] = true
		knots[i] = Knot{Point{0, 0}, Point{0, 0}, visited, &knots[i-1]}
	}
	return knots
}
