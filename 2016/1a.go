package main

import (
	"fmt"
	"image"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2016
// Day 1, part a

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
	fmt.Println(algo.Abs(p.X) + algo.Abs(p.Y))
}
