package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Push(v string) {
	*s = append(*s, v)
}

// PushBottom - Push to bottom of the stack
func (s *Stack) PushBottom(v string) {
	*s = append([]string{v}, *s...)
}

func (s *Stack) Pop() (string, bool) {
	l := len(*s)
	if l == 0 {
		return "", false
	}
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res, true
}

type Instruction struct {
	Number    int
	FromIndex int
	ToIndex   int
}

func NewInstruction(line string) (*Instruction, error) {
	words := strings.Split(line, " ")
	if len(words) >= 5 {
		amount, err := strconv.Atoi(words[1])
		fromNum, err := strconv.Atoi(words[3])
		toNum, err := strconv.Atoi(words[5])

		if err != nil {
			panic(err)
		}
		fromIndex := fromNum - 1
		toIndex := toNum - 1

		return &Instruction{
			Number:    amount,
			FromIndex: fromIndex,
			ToIndex:   toIndex,
		}, nil
	}
	return nil, fmt.Errorf("invalid instruction: %s", line)
}
