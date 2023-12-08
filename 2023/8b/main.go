package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 8, part b

const inputPath = "2023/inputs/8.txt"

func main() {
	lines := exp.MustReadLines(inputPath)

	type node struct {
		l, r string
	}
	m := make(map[string]node)

	var pos, targ []string
	inst := lines[0]
	for _, line := range lines[2:] {
		m[line[0:3]] = node{l: line[7:10], r: line[12:15]}
		if line[2] == 'A' {
			pos = append(pos, line[0:3])
		}
		if line[2] == 'Z' {
			targ = append(targ, line[0:3])
		}
	}

	//fmt.Println(pos)

	type state struct {
		step int
		pos  string
	}

	seen := make([]map[state]int, len(pos))
	cyc := make([]int, len(pos))

	for i := range pos {
		step := 0
		seen[i] = map[state]int{
			{step: 0, pos: pos[i]}: 0,
		}
		for {
			switch inst[step%len(inst)] {
			case 'L':
				pos[i] = m[pos[i]].l
			case 'R':
				pos[i] = m[pos[i]].r
			}
			step++

			st := state{step: step % len(inst), pos: pos[i]}
			if last, y := seen[i][st]; y {
				cyc[i] = step - last
				break
			}
			seen[i][st] = step
		}
	}

	n := 1
	for _, c := range cyc {
		d := algo.GCD(n, c)
		c /= d
		n *= c
	}

	fmt.Println(n)
}
