package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2015
// Day 16, part a

func main() {
	want := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	for _, line := range exp.MustReadLines("inputs/16.txt") {
		var num int
		pre, post, ok := strings.Cut(line, ": ")
		if !ok {
			log.Fatalf("Couldn't cut line %q", line)
		}
		exp.Must(fmt.Sscanf(pre, "Sue %d", &num))

		match := true
		for _, att := range strings.Split(post, ", ") {
			k, v, ok := strings.Cut(att, ": ")
			if !ok {
				log.Fatalf("Couldn't cut attribute %q", line)
			}
			if got := exp.Must(strconv.Atoi(v)); got != want[k] {
				match = false
				break
			}
		}

		if match {
			fmt.Println(num)
		}
	}
}
