package d11_lib

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

func (ms *MonkeySystem) ProcessMonkeyItemsPart1(m *Monkey) {
	item := m.OperateOnNextItemPart1()
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
		ms.ProcessMonkeyItemsPart1(m)
	}
}

func (ms *MonkeySystem) ProcessMonkeysPart1() {
	for _, m := range ms.Monkeys {
		ms.ProcessMonkeyItemsPart1(m)
	}
}
