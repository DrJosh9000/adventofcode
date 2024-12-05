package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2023
// Day 20, part b

const inputPath = "2023/inputs/20.txt"

type module interface {
	outs() []string
}

type broadcaster struct {
	out []string
}

type flipflop struct {
	out []string
}

type conjunction struct {
	out []string
}

func (m *broadcaster) outs() []string { return m.out }
func (m *flipflop) outs() []string    { return m.out }
func (m *conjunction) outs() []string { return m.out }

func main() {
	lines := exp.MustReadLines(inputPath)
	byName := make(map[string]module)
	for _, line := range lines {
		in, out, ok := strings.Cut(line, " -> ")
		if !ok {
			panic("invalid line " + line)
		}
		outs := strings.Split(out, ", ")
		var m module
		var name string
		switch {
		case in[0] == '%':
			name = in[1:]
			m = &flipflop{
				out: outs,
			}

		case in[0] == '&':
			name = in[1:]
			m = &conjunction{
				out: outs,
			}

		case in == "broadcaster":
			name = in
			m = &broadcaster{
				out: outs,
			}

		default:
			panic(fmt.Sprintf("invalid name %q", name))
		}

		byName[name] = m
	}

	var highs []int

	// (some scribbling on paper)
	// these flipflops and conjunctions are wired suspiciously like four
	// pulse counters

	for _, mod := range byName["broadcaster"].outs() {
		ff, ok := byName[mod].(*flipflop)
		if !ok {
			panic("not a flipflop")
		}
		n, b := 0, 1
		chain := true
		for chain {
			chain = false
			for _, ffo := range ff.outs() {
				switch next := byName[ffo].(type) {
				case *conjunction:
					n += b
				case *flipflop:
					ff = next
					chain = true
				}
			}
			b *= 2
		}
		highs = append(highs, n)
	}

	//fmt.Println(highs)

	n := 1
	for _, h := range highs {
		d := algo.GCD(n, h)
		n /= d
		n *= h
	}

	fmt.Println(n)
}
