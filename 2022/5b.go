package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2022
// Day 5, part b

func main() {
	// top of each stack is at 0
	var stacks [][]byte

	movesMode := false
	for _, line := range exp.MustReadLines("inputs/5.txt") {
		if movesMode {
			var q, s, d int
			exp.Must(fmt.Sscanf(line, "move %d from %d to %d", &q, &s, &d))
			s--
			d--

			t := make([]byte, q)
			copy(t, stacks[s][:q])
			stacks[s] = stacks[s][q:]
			stacks[d] = append(t, stacks[d]...)
			continue
		}
		if line == "" {
			movesMode = true
			continue
		}
		n := (len(line) + 1) / 4
		if len(stacks) < n {
			stacks = append(stacks, make([][]byte, n-len(stacks))...)
		}
		for i := 0; i < n; i++ {
			c := line[i*4+1]
			if c == ' ' {
				continue
			}
			stacks[i] = append(stacks[i], c)
		}
	}

	for i := range stacks {
		fmt.Printf("%c", stacks[i][0])
	}
	fmt.Println()
}
