package d11_p1

import (
	lib "day_11/d11_lib"
	"fmt"
)

const Rounds = 20

func Part1(lines []string) {
	fmt.Println("Day 11, Part 1")
	monkeySystem := lib.BuildMonkeySystem(lines)

	for i := 0; i < Rounds; i++ {
		monkeySystem.ProcessMonkeysPart1()
	}

	highestNum := -1
	highestNum2 := -1
	for _, m := range monkeySystem.Monkeys {
		// get the two highest numbers from all m.NumItemsInspected
		if m.NumItemsInspected > highestNum {
			highestNum2 = highestNum
			highestNum = m.NumItemsInspected
		} else if m.NumItemsInspected > highestNum2 {
			highestNum2 = m.NumItemsInspected
		}
	}
	fmt.Println("Highest Num: ", highestNum)
	fmt.Println("Highest Num2: ", highestNum2)
	fmt.Println("Product of the two highest numbers: ", highestNum*highestNum2)
}
