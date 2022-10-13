package main

import (
	"fmt"
	"regexp"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2015
// Day 5, part a

var subRE = regexp.MustCompile(`ab|cd|pq|xy`)

func main() {
	nice := 0
	for _, line := range exp.MustReadLines("inputs/5.txt") {
		if subRE.MatchString(line) {
			continue
		}
		h := algo.Freq([]byte(line))
		if v := h['a'] + h['e'] + h['i'] + h['o'] + h['u']; v < 3 {
			continue
		}
		d := false
		for i := range line[1:] {
			if line[i] == line[i+1] {
				d = true
				break
			}
		}
		if !d {
			continue
		}
		nice++
	}
	fmt.Println(nice)
}
