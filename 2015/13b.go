package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2015
// Day 13, part b

func main() {
	input := exp.MustReadLines("inputs/13.txt")

	label := make(map[string]int)
	for _, line := range input {
		fs := strings.Fields(strings.TrimSuffix(line, "."))
		if _, old := label[fs[0]]; !old {
			label[fs[0]] = len(label)
		}
		if _, old := label[fs[len(fs)-1]]; !old {
			label[fs[len(fs)-1]] = len(label)
		}
	}

	label["Myself"] = len(label)

	N := len(label)
	dh := grid.Make[int](N, N) // x would gain dh[x][y] units by sitting next to y
	for _, line := range input {
		fs := strings.Fields(strings.TrimSuffix(line, "."))
		a := exp.Must(strconv.Atoi(fs[3]))
		s, t := label[fs[0]], label[fs[len(fs)-1]]
		if fs[2] == "lose" {
			a *= -1
		}
		dh[s][t] = a
	}

	ord := make([]int, N)
	for i := range ord {
		ord[i] = i
	}

	maxh := math.MinInt
	for {
		h := 0
		for i, p := range ord {
			h += dh[p][ord[(i+N-1)%N]]
			h += dh[p][ord[(i+1)%N]]
		}
		if h > maxh {
			maxh = h
		}

		if !algo.NextPermutation(ord) {
			break
		}
	}

	fmt.Println(maxh)
}
