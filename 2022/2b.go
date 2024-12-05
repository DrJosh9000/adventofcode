package main

import (
	"fmt"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2022
// Day 2, part b

var shape = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var choice = map[string]string{
	"A X": "C", // rock lose = scissors
	"A Y": "A", // rock draw = rock
	"A Z": "B", // rock win = paper
	"B X": "A", // paper lose = rock
	"B Y": "B", // paper draw = paper
	"B Z": "C", // paper win = scissors
	"C X": "B", // scissors lose = paper
	"C Y": "C", // scissors draw = scissors
	"C Z": "A", // scissors win = rock
}

var outcome = map[string]int{
	"A A": 3, // rock vs rock
	"A B": 6, // rock vs paper
	"A C": 0, // rock vs scissors
	"B A": 0, // paper vs rock
	"B B": 3, // paper vs paper
	"B C": 6, // paper vs scissors
	"C A": 6, // scissors vs rock
	"C B": 0, // scissors vs paper
	"C C": 3, // scissors vs scissors
}

func main() {
	score := 0
	for _, line := range exp.MustReadLines("inputs/2.txt") {
		m := strings.Fields(line)
		c := choice[line]
		score += outcome[m[0]+" "+c]
		score += shape[c]
	}
	fmt.Println(score)
}
