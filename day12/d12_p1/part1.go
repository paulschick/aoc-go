package d12_p1

import "fmt"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

type AlphabetMap struct {
	LetterToNumber map[rune]int
}

func NewAlphabetMap() *AlphabetMap {
	letterToNumber := make(map[rune]int)
	for i, letter := range alphabet {
		letterToNumber[letter] = i
	}
	return &AlphabetMap{LetterToNumber: letterToNumber}
}

func (a *AlphabetMap) MapLetter(c rune) int {
	return a.LetterToNumber[c]
}

func Part1(lines []string) {
	fmt.Println("Day 12, Part 1")
	for _, line := range lines {
		fmt.Println(line)
	}
}
