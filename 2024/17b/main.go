package main

import (
	"fmt"
	"os"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2024
// Day 17, part b

const inputPath = "2024/inputs/17.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	var prog []int
	for _, line := range lines {
		if s := ""; exp.Smatchf(line, "Program: %s\n", &s) {
			prog = algo.Map(strings.Split(s, ","), exp.MustAtoi)
		}
	}

	// fmt.Printf("Registers: %v\n", reg)
	// fmt.Printf("Program: %v\n", prog)

	var search func(int, int)
	search = func(a, d int) {
		if d < 0 {
			fmt.Println(a)
			os.Exit(0)
		}
		want := prog[d]
		a <<= 3
		for n := range 0o10 {
			t := a + n
			got := run(prog, t)
			if got != want {
				continue
			}
			search(t, d-1)
		}
	}

	search(0, len(prog)-1)
}

// 2,4,  B = A % 8
// 1,1,  B ^= 1
// 7,5,  C = A >> B
// 1,5,  B ^= 5
// 4,0,  B ^= C
// 5,5,  out B % 8
// 0,3,  A >>= 3
// 3,0   if A != 0 { goto 0 }

func run(prog []int, a int) int {
	var reg [3]int
	// out := make([]int, 0, len(prog))
	reg[0] = a
	ip := 0
	literal := func() int { return prog[ip+1] }
	combo := func() int {
		x := prog[ip+1]
		if x < 4 {
			return x
		}
		return reg[x-4]
	}
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
			return combo() & 0x7

		case 6: // bdv
			reg[1] = reg[0] >> combo()

		case 7: // cdv
			reg[2] = reg[0] >> combo()
		}
		ip += 2
	}
	return -1
}
