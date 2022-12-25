package main

import (
	"fmt"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2022
// Day 25

func snafu(x int) string {
	if x == 0 {
		return "0"
	}
	drev := map[int]rune{
		0:  '0',
		1:  '1',
		2:  '2',
		-1: '-',
		-2: '=',
	}

	var b []rune
	for x != 0 {
		d := x % 5
		if d > 2 {
			// must be expressed as a negative
			x += 5
			d -= 5
		}
		b = append(b, drev[d])
		x /= 5
	}
	algo.Reverse(b)
	return string(b)
}

func main() {
	//fmt.Println(314159265, snafu(314159265))

	digits := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'-': -1,
		'=': -2,
	}

	sum := 0
	for _, line := range exp.MustReadLines("inputs/25.txt") {
		x := 0
		for _, c := range line {
			x *= 5
			x += digits[c]
		}
		sum += x
	}

	fmt.Println("In deicmal:", sum)

	fmt.Println(snafu(sum))
}
