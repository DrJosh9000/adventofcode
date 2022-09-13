package main

import (
	"fmt"
	"image"
	"strconv"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2016
// Day 1, part a

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	p, d := image.Pt(0, 0), image.Pt(0, -1)
	for _, step := range exp.MustReadDelimited("inputs/1.txt", ",") {
		if step[0] == 'L' {
			d = image.Pt(d.Y, -d.X)
		} else {
			d = image.Pt(-d.Y, d.X)
		}
		p = p.Add(d.Mul(exp.Must(strconv.Atoi(step[1:]))))
	}
	fmt.Println(abs(p.X) + abs(p.Y))
}
