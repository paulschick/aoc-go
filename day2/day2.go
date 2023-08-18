package main

import (
	"aoc/input_data_reader"
	"fmt"
)

/*
*
Part 1
Col 1 - opponent
A - rock
B - paper
C - scissors

Col 2 - you
X - rock
Y - paper
Z - scissors

Shape score per round:
X - 1
Y - 2
Z - 3

Outcome score per round:
lose - 0
draw - 3
win - 6

total per round - shape score + outcome score
Get the total score as the sum of all rounds
*/
func part1(lines []string) {
	totalScore := 0

	for _, line := range lines {
		roundScore := 0

		opponent := []rune(line)[0]
		me := []rune(line)[2]

		switch opponent {
		case 'A':
			switch me {
			case 'X':
				// draw
				roundScore = 4
			case 'Y':
				// win
				roundScore = 8
			case 'Z':
				// lose
				roundScore = 3
			}
		case 'B':
			switch me {
			case 'X':
				// lose
				roundScore = 1
			case 'Y':
				// draw
				roundScore = 5
			case 'Z':
				// win
				roundScore = 9
			}
		case 'C':
			switch me {
			case 'X':
				// win
				roundScore = 7
			case 'Y':
				// lose
				roundScore = 2
			case 'Z':
				// draw
				roundScore = 6
			}
		}
		totalScore += roundScore
	}

	fmt.Println("---------- Part 1 Final Score ----------")
	fmt.Println("total score: ", totalScore)
}

/*
*
part 2
X - lose
Y - draw
Z - win
*/
func part2(lines []string) {
	totalScore := 0

	for _, line := range lines {
		roundScore := 0

		opponent := []rune(line)[0]
		me := []rune(line)[2]

		switch opponent {
		// rock
		case 'A':
			switch me {
			case 'X':
				// need to lose
				// choose scissors - 3 + 0
				roundScore = 3
			case 'Y':
				// need to draw
				// choose rock - 1 + 3
				roundScore = 4
			case 'Z':
				// need to win
				// choose paper - 2 + 6
				roundScore = 8
			}
		// paper
		case 'B':
			switch me {
			case 'X':
				// need to lose
				// choose rock - 1 + 0
				roundScore = 1
			case 'Y':
				// need to draw
				// choose paper - 2 + 3
				roundScore = 5
			case 'Z':
				// need to win
				// choose scissors - 3 + 6
				roundScore = 9
			}
		// scissors
		case 'C':
			switch me {
			case 'X':
				// need to lose
				// choose paper - 2 + 0
				roundScore = 2
			case 'Y':
				// need to draw
				// choose scissors - 3 + 3
				roundScore = 6
			case 'Z':
				// need to win
				// choose rock - 1 + 6
				roundScore = 7
			}
		}
		totalScore += roundScore
	}

	fmt.Println("---------- Part 2 Final Score ----------")
	fmt.Println("total score: ", totalScore)
}

func main() {
	lines := input_data_reader.Reader("day2.txt")

	fmt.Println("day 2, part 1")
	part1(lines)
	fmt.Println("day 2, part 2")
	part2(lines)
}
