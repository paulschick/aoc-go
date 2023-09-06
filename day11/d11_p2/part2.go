package d11_p2

import (
	"aoc/utils"
	lib "day_11/d11_lib"
	"fmt"
)

const rounds = 10_000

func Part2(lines []string) {
	fmt.Println("Day 11, Part 2")
	monkeySystem := lib.BuildMonkeySystem(lines)

	commonDivisor := 1
	for _, m := range monkeySystem.Monkeys {
		commonDivisor *= m.TestDivisor
	}

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeySystem.Monkeys {
			for _, item := range monkey.Items {
				monkey.NumItemsInspected++
				newLevel := monkey.Operation(item)
				newLevel %= commonDivisor

				if newLevel%monkey.TestDivisor == 0 {
					trueMonkey := monkeySystem.MonkeyLookup[monkeySystem.RoutingRules[monkey.Number][true]]
					trueMonkey.SetItems(append(trueMonkey.Items, newLevel))
				} else {
					falseMonkey := monkeySystem.MonkeyLookup[monkeySystem.RoutingRules[monkey.Number][false]]
					falseMonkey.SetItems(append(falseMonkey.Items, newLevel))
				}
			}
			monkey.SetItems(make([]int, 0))
		}
	}

	inspectedCounts := make([]int, len(monkeySystem.Monkeys))
	for i, m := range monkeySystem.Monkeys {
		inspectedCounts[i] = m.NumItemsInspected
	}

	// sort highest to lowest
	utils.SortInts(inspectedCounts)
	highest := inspectedCounts[0]
	highest2 := inspectedCounts[1]
	fmt.Println("Highest Num: ", highest)
	fmt.Println("Highest Num2: ", highest2)
	mult := highest * highest2
	fmt.Println("Product of the two highest numbers: ", mult)
}
