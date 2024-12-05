package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2023
// Day 8, part a

const inputPath = "2023/inputs/8.txt"

func main() {
	lines := exp.MustReadLines(inputPath)

	type node struct {
		l, r string
	}
	m := make(map[string]node)

	inst := lines[0]
	for _, line := range lines[2:] {
		m[line[0:3]] = node{l: line[7:10], r: line[12:15]}
	}
	//fmt.Println(inst, m)

	step := 0
	pos := "AAA"
	for pos != "ZZZ" {
		switch inst[step%len(inst)] {
		case 'L':
			pos = m[pos].l
		case 'R':
			pos = m[pos].r
		}
		step++
	}
	fmt.Println(step)
}
