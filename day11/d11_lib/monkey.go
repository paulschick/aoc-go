package d11_lib

import (
	"strconv"
	"strings"
)

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
	TestDivisor       int
}

func NewMonkey() *Monkey {
	return &Monkey{}
}

// OperateOnNextItemPart1 pops the front item from the queue slice.
// Runs the operation and then divides by 3.
func (m *Monkey) OperateOnNextItemPart1() int {
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

func (m *Monkey) SetTestDivisor(value int) {
	m.TestDivisor = value
}

func GetCurrentMonkeyNumber(stringFields []string) int {
	monkeyNumStr := strings.Split(stringFields[1], ":")[0]
	currentMonkeyNumber, err := strconv.Atoi(monkeyNumStr)
	if err != nil {
		panic(err)
	}
	return currentMonkeyNumber
}
