package day10_part2

import (
	"fmt"
	"strconv"
)

const PixelWidth = 40

func Part2(lines []string) {
	fmt.Println("Day 10, Part 2")

	screen := NewScreen()

	for _, line := range lines {
		switch line[:4] {
		case "addx":
			addX, err := strconv.Atoi(line[5:])
			if err == nil {
				screen.AddX(addX)
			}
		default:
			screen.RenderOneCycle()
		}
	}

	for _, row := range screen.RenderedRows {
		fmt.Println(row)
	}
}

type Screen struct {
	X            int
	Cycle        int
	CurrentRow   string
	RenderedRows []string
}

func NewScreen() *Screen {
	return &Screen{X: 1, Cycle: -1}
}

func (s *Screen) RenderOneCycle() {
	s.IncrementOneCycle()
	s.AppendPixelToRow()
	s.CheckAndAppendRow()
}

func (s *Screen) IncrementOneCycle() {
	s.Cycle++
}

func (s *Screen) AppendPixelToRow() {
	currentPixel := s.CurrentPixel()

	if s.IsWithinSpriteRange(currentPixel) {
		s.CurrentRow += "#"
	} else {
		s.CurrentRow += "."
	}
}

func (s *Screen) CheckAndAppendRow() {
	if s.Cycle%PixelWidth == PixelWidth-1 {
		s.RenderedRows = append(s.RenderedRows, s.CurrentRow)
		s.CurrentRow = ""
	}
}

func (s *Screen) IsWithinSpriteRange(pixel int) bool {
	spriteLeft := s.X - 1
	spriteRight := s.X + 1
	return spriteLeft <= pixel && pixel <= spriteRight
}

func (s *Screen) AddX(value int) {
	s.RenderOneCycle()
	s.RenderOneCycle()
	s.X += value
}

func (s *Screen) CurrentPixel() int {
	return s.Cycle % PixelWidth
}
