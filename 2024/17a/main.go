package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2024
// Day 17, part a

const inputPath = "2024/inputs/17.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	var reg [3]int
	var prog []int
	for _, line := range lines {
		var r rune
		var x int
		var s string
		switch {
		case exp.Smatchf(line, "Register %c: %d", &r, &x):
			reg[r-'A'] = x
		case exp.Smatchf(line, "Program: %s\n", &s):
			prog = algo.Map(strings.Split(s, ","), exp.MustAtoi)
		}
	}

	// fmt.Printf("Registers: %v\n", reg)
	// fmt.Printf("Program: %v\n", prog)

	ip := 0
	literal := func() int { return prog[ip+1] }
	combo := func() int {
		x := prog[ip+1]
		if x < 4 {
			return x
		}
		return reg[x-4]
	}

	comma := false
	for ip >= 0 && ip < len(prog) {
		switch prog[ip] {
		case 0: // adv
			reg[0] >>= combo()

		case 1: // bxl
			reg[1] ^= literal()

		case 2: // bst
			reg[1] = combo() & 0x7

		case 3: // jnz
			if reg[0] != 0 {
				ip = literal()
				continue
			}

		case 4: // bxc
			reg[1] ^= reg[2]

		case 5:
			if comma {
				fmt.Print(",")
			}
			comma = true
			fmt.Print(combo() & 0x7)

		case 6: // bdv
			reg[1] = reg[0] >> combo()

		case 7: // cdv
			reg[2] = reg[0] >> combo()
		}
		ip += 2
	}
	fmt.Println()
}
