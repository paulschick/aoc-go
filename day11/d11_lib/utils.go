package d11_lib

import (
	"strconv"
	"strings"
)

func BuildVariableOperation(operator string, variableValue int) func(i int) int {
	return func(i int) int {
		switch operator {
		case "+":
			return i + variableValue
		case "-":
			return i - variableValue
		case "*":
			return i * variableValue
		case "/":
			return i / variableValue
		default:
			panic("Unknown operator")
		}
	}
}

func BuildOldOperation(operator string) func(i int) int {
	return func(i int) int {
		switch operator {
		case "+":
			return i + i
		case "-":
			return i - i
		case "*":
			return i * i
		case "/":
			return i / i
		default:
			panic("Unknown operator")
		}
	}
}

func GetTestFunc(divisibleByInt int) func(i int) bool {
	return func(i int) bool {
		return i%divisibleByInt == 0
	}
}

func BuildMonkeySystem(lines []string) *MonkeySystem {
	var currentMonkey *Monkey
	monkeySystem := NewMonkeySystem()

	for _, line := range lines {
		values := strings.Fields(line)
		if len(values) > 0 {
			switch values[0] {
			case "Monkey":
				currentMonkey = NewMonkey()
				currentMonkeyNumber := GetCurrentMonkeyNumber(values)
				currentMonkey.SetNumber(currentMonkeyNumber)
				monkeySystem.AddMonkey(currentMonkey)

			case "Starting":
				newItems := values[2:]
				newItemsInt := make([]int, 0)
				for _, v := range newItems {
					numberSplit := strings.Split(v, ",")
					itemInt, err := strconv.Atoi(numberSplit[0])
					if err != nil {
						panic(err)
					}
					newItemsInt = append(newItemsInt, itemInt)
				}
				currentMonkey.SetItems(newItemsInt)
			case "Operation:":
				operator := values[4]
				variable := values[5]

				if variable != "old" {
					variableValue, err := strconv.Atoi(variable)
					if err != nil {
						panic(err)
					}
					currentMonkey.SetOperation(BuildVariableOperation(operator, variableValue))
				} else {
					currentMonkey.SetOperation(BuildOldOperation(operator))
				}
			case "Test:":
				divisibleBy := values[3]
				divisibleByInt, err := strconv.Atoi(divisibleBy)
				if err != nil {
					panic(err)
				}
				currentMonkey.SetTest(GetTestFunc(divisibleByInt))
				currentMonkey.SetTestDivisor(divisibleByInt)
			case "If":
				toMonkey, err := strconv.Atoi(values[5])
				if err != nil {
					panic(err)
				}
				if values[1] == "true:" {
					monkeySystem.SetRoutingRule(currentMonkey.Number, true, toMonkey)
				} else {
					monkeySystem.SetRoutingRule(currentMonkey.Number, false, toMonkey)
				}
			}
		}
	}
	return monkeySystem
}
