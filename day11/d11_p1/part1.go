package d11_p1

import (
	"fmt"
	"strconv"
	"strings"
)

const Rounds = 20

func Part1(lines []string) {
	fmt.Println("Day 11, Part 1")
	monkeySystem := BuildMonkeySystem(lines)

	for i := 0; i < Rounds; i++ {
		monkeySystem.ProcessMonkeys()
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

// Monkey Represents a set of items and operations
/**
 * Number: The number of the monkey
 * Items: The items to be operated on
 * Operation: The operation to be performed on the items
 * Test: The test to determine which monkey index to send to
 * SendTrueIndex: The index of the monkey to send to if test is true
 * SendFalseIndex: The index of the monkey to send to if test is false
 *
 * When working on an item, first do the Operation. Next, divide by 3, and after the division,
 * do the Test. If the Test is true, send the item to the SendTrueIndex monkey. If the Test is
 * false, send the item to the SendFalseIndex monkey.
 */
type Monkey struct {
	Number            int
	Items             []int
	Operation         func(int) int
	Test              func(int) bool
	NumItemsInspected int
}

func NewMonkey() *Monkey {
	return &Monkey{}
}

// OperateOnNextItem pops the front item from the queue slice.
// Runs the operation and then divides by 3.
func (m *Monkey) OperateOnNextItem() int {
	if (len(m.Items)) == 0 {
		return -1
	}
	oldValue := m.Items[0]
	m.SetItems(m.Items[1:])
	m.NumItemsInspected++
	return m.Operation(oldValue) / 3
}

func (m *Monkey) SetNumber(number int) {
	m.Number = number
}

func (m *Monkey) SetItems(items []int) {
	m.Items = items
}

func (m *Monkey) SetOperation(operation func(int) int) {
	m.Operation = operation
}

func (m *Monkey) SetTest(test func(int) bool) {
	m.Test = test
}

func GetCurrentMonkeyNumber(stringFields []string) int {
	monkeyNumStr := strings.Split(stringFields[1], ":")[0]
	currentMonkeyNumber, err := strconv.Atoi(monkeyNumStr)
	if err != nil {
		panic(err)
	}
	return currentMonkeyNumber
}

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

// MonkeySystem represents a system of monkeys to manage the collection and routing of items.
type MonkeySystem struct {
	Monkeys      []*Monkey
	MonkeyLookup map[int]*Monkey
	RoutingRules map[int]map[bool]int
}

func NewMonkeySystem() *MonkeySystem {
	return &MonkeySystem{
		MonkeyLookup: make(map[int]*Monkey),
		RoutingRules: make(map[int]map[bool]int),
	}
}

func (ms *MonkeySystem) AddMonkey(m *Monkey) {
	ms.Monkeys = append(ms.Monkeys, m)
	ms.MonkeyLookup[m.Number] = m
}

// SetRoutingRule sets the routing rule for a monkey.
// This defines which Monkey to send an item to when an outcome is true or false.
func (ms *MonkeySystem) SetRoutingRule(monkeyId int, outcome bool, targetId int) {
	if _, ok := ms.RoutingRules[monkeyId]; !ok {
		ms.RoutingRules[monkeyId] = make(map[bool]int)
	}
	ms.RoutingRules[monkeyId][outcome] = targetId
}

func (ms *MonkeySystem) ProcessMonkeyItems(m *Monkey) {
	item := m.OperateOnNextItem()
	if item == -1 {
		return
	}
	// already divided by three, so run test
	testValue := m.Test(item)
	if testValue {
		// Grab the monkey's items for true condition
		trueMonkey := ms.MonkeyLookup[ms.RoutingRules[m.Number][true]]
		trueMonkey.SetItems(append(trueMonkey.Items, item))
	} else {
		// Grab the monkey's items for false condition
		falseMonkey := ms.MonkeyLookup[ms.RoutingRules[m.Number][false]]
		falseMonkey.SetItems(append(falseMonkey.Items, item))
	}

	if len(m.Items) > 0 {
		ms.ProcessMonkeyItems(m)
	}
}

func (ms *MonkeySystem) ProcessMonkeys() {
	for _, m := range ms.Monkeys {
		ms.ProcessMonkeyItems(m)
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
