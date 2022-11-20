package main

import (
	"fmt"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2015
// Day 19, part a

func main() {
	repl := make(map[string][]string)
	var start string
	for _, line := range exp.MustReadLines("inputs/19.txt") {
		if line == "" {
			continue
		}
		in, out, ok := strings.Cut(line, " => ")
		if !ok {
			start = line
			continue
		}
		repl[in] = append(repl[in], out)
	}

	mol := make(algo.Set[string])
	for i := range start {
		for in, out := range repl {
			if strings.HasPrefix(start[i:], in) {
				for _, o := range out {
					mol.Insert(start[:i] + o + start[i+len(in):])
				}
			}
		}
	}
	fmt.Println(len(mol))
}
