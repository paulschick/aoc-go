package day10_part1

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Part1(lines []string) {
	fmt.Println("Day 10, Part 1")
	instructions := ParseInstructions(lines)
	cycleNumber := 0
	cycleIncrement := CycleIncrement{0, 1, 1, 1}
	part1Cycles := Part1Cycles{0, 0, 0, 0, 0, 0}

	for _, instruction := range instructions {
		numCycles := instruction.Cycles
		xMove := instruction.XMove

		for i := 0; i < numCycles; i++ {
			cycleNumber++
			cycleIncrement.Cycle = cycleNumber

			cycleMove := 0
			if i > 0 {
				cycleMove = xMove
			}
			cycleInstruction := NewCycleInstruction(cycleNumber, cycleMove)
			fmt.Println(cycleInstruction.String())

			cycleIncrement.XStart = cycleIncrement.XEnd
			cycleIncrement.XSignalStrength = cycleIncrement.XStart * cycleInstruction.XMult
			cycleIncrement.XEnd = cycleIncrement.XStart + cycleMove
			fmt.Println(cycleIncrement.String())

			if cycleNumber%20 == 0 {
				part1Cycles = *part1Cycles.AddValue(cycleNumber, cycleIncrement.XSignalStrength)
				fmt.Println(part1Cycles.String())
			}
		}
	}

	fmt.Println("Part 1 Cycles")
	fmt.Println(part1Cycles.SumCycles())
}

// Part1Cycles is a struct that holds the cycles to be summed for Part 1
// This refers to the value of X at each of these cycles
type Part1Cycles struct {
	Cycle20  int
	Cycle60  int
	Cycle100 int
	Cycle140 int
	Cycle180 int
	Cycle220 int
}

func (p1c *Part1Cycles) AddValue(cycleNumber int, signalStrength int) *Part1Cycles {
	switch {
	case cycleNumber == 20:
		p1c.Cycle20 = signalStrength
	case cycleNumber == 60:
		p1c.Cycle60 = signalStrength
	case cycleNumber == 100:
		p1c.Cycle100 = signalStrength
	case cycleNumber == 140:
		p1c.Cycle140 = signalStrength
	case cycleNumber == 180:
		p1c.Cycle180 = signalStrength
	case cycleNumber == 220:
		p1c.Cycle220 = signalStrength
	}
	return p1c
}

func (p1c *Part1Cycles) String() string {
	return fmt.Sprintf("Cycle20: %d, Cycle60: %d, Cycle100: %d, Cycle140: %d, Cycle180: %d, Cycle220: %d", p1c.Cycle20, p1c.Cycle60, p1c.Cycle100, p1c.Cycle140, p1c.Cycle180, p1c.Cycle220)
}

// SumCycles Returns the sum of all the cycles for Part 1
func (p1c *Part1Cycles) SumCycles() int {
	return p1c.Cycle20 + p1c.Cycle60 + p1c.Cycle100 + p1c.Cycle140 + p1c.Cycle180 + p1c.Cycle220
}

// CycleIncrement is a struct that holds the cycle and the XStart and XEnd values
// XStart is equivalent to the ending value of the previous cycle
// XEnd is the final value of X for that cycle
type CycleIncrement struct {
	Cycle           int
	XStart          int
	XSignalStrength int
	XEnd            int
}

func (ci *CycleIncrement) String() string {
	return fmt.Sprintf("Cycle: %d, XStart: %d, XSignalStrength: %d, XEnd: %d", ci.Cycle, ci.XStart, ci.XSignalStrength, ci.XEnd)
}

// CycleInstruction is a struct that holds the cycle and the XMove and XMult values
// This represents a single cycle within an instruction
// for a noop, Cycle is the cycle number, XMove is 0, and if Cycle is a multiplication cycle,
// then XMult would be equal to the cycle number
// CycleInstruction would be used to get the value CycleIncrement for that cycle.
// XMult can just default to 1 so that the behavior can be the same for each of the cycles
type CycleInstruction struct {
	Cycle int
	XMove int
	XMult int
}

func (ci *CycleInstruction) String() string {
	return fmt.Sprintf("Cycle: %d, XMove: %d, XMult: %d", ci.Cycle, ci.XMove, ci.XMult)
}

func NewCycleInstruction(cycle int, xMove int) *CycleInstruction {
	xMult := 1
	if cycle%20 == 0 {
		xMult = cycle
	}
	return &CycleInstruction{cycle, xMove, xMult}
}

// Instruction is a struct that holds the cycles and the XMove value
// NOOP has 1 cycle and XMove is 0
type Instruction struct {
	Cycles int
	XMove  int
}

func ParseInstructions(input []string) []Instruction {
	instructions := make([]Instruction, len(input))
	for i, line := range input {
		if line == "noop" {
			// add noop instruction
			instructions[i] = Instruction{Cycles: 1, XMove: 0}
		} else {
			// add addx instruction
			splitLine := strings.Split(line, " ")
			moveX := utils.ToInt(splitLine[1])
			instructions[i] = Instruction{Cycles: 2, XMove: moveX}
		}
	}
	return instructions
}
