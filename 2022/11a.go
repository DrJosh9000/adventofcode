package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 11, part a

type monkey struct {
	num    int
	items  []int
	op     rune
	opB    any
	mod    int
	dtrue  int
	dfalse int
	insp   int
}

func (m *monkey) worry(x int) int {
	y, ok := m.opB.(int)
	if !ok {
		y = x
	}
	switch m.op {
	case '+':
		return (x + y) / 3
	case '*':
		return (x * y) / 3
	}
	panic(fmt.Sprintf("unknown operator %c", m.op))
}

func main() {
	var monkeys []*monkey
	var m *monkey
	for _, line := range exp.MustReadLines("inputs/11.txt") {
		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "Monkey"):
			m = new(monkey)
			monkeys = append(monkeys, m)
			exp.Must(fmt.Sscanf(line, "Monkey %d:", &m.num))
		case strings.HasPrefix(line, "Starting items:"):
			line = strings.TrimPrefix(line, "Starting items: ")
			for _, i := range strings.Split(line, ", ") {
				m.items = append(m.items, exp.Must(strconv.Atoi(i)))
			}
		case strings.HasPrefix(line, "Operation:"):
			var opb string
			exp.Must(fmt.Sscanf(line, "Operation: new = old %c %s", &m.op, &opb))
			n, err := strconv.Atoi(opb)
			if err != nil {
				m.opB = opb
			} else {
				m.opB = n
			}
		case strings.HasPrefix(line, "Test:"):
			exp.Must(fmt.Sscanf(line, "Test: divisible by %d", &m.mod))
		case strings.HasPrefix(line, "If true:"):
			exp.Must(fmt.Sscanf(line, "If true: throw to monkey %d", &m.dtrue))
		case strings.HasPrefix(line, "If false:"):
			exp.Must(fmt.Sscanf(line, "If false: throw to monkey %d", &m.dfalse))
		}
	}

	for round := 0; round < 20; round++ {
		for _, m := range monkeys {
			for _, i := range m.items {
				j := m.worry(i)
				if j%m.mod == 0 {
					monkeys[m.dtrue].items = append(monkeys[m.dtrue].items, j)
				} else {
					monkeys[m.dfalse].items = append(monkeys[m.dfalse].items, j)
				}
			}
			m.insp += len(m.items)
			m.items = nil
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].insp > monkeys[j].insp
	})
	fmt.Println(monkeys[0].insp * monkeys[1].insp)
}
