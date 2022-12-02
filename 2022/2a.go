package main

import (
	"fmt"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2022
// Day 2, part a

var shape = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var outcome = map[string]int{
	"A X": 3, // rock vs rock
	"A Y": 6, // rock vs paper
	"A Z": 0, // rock vs scissors
	"B X": 0, // paper vs rock
	"B Y": 3, // paper vs paper
	"B Z": 6, // paper vs scissors
	"C X": 6, // scissors vs rock
	"C Y": 0, // scissors vs paper
	"C Z": 3, // scissors vs scissors
}

func main() {
	score := 0
	for _, line := range exp.MustReadLines("inputs/2.txt") {
		m := strings.Fields(line)
		score += outcome[line]
		score += shape[m[1]]
	}
	fmt.Println(score)
}
