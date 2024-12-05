package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 16, part b

func eq(a, b int) bool { return a == b }
func gt(a, b int) bool { return a > b }
func lt(a, b int) bool { return a < b }

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
	matchFunc := map[string]func(int, int) bool{
		"cats":        gt,
		"trees":       gt,
		"pomeranians": lt,
		"goldfish":    lt,
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
			mf := matchFunc[k]
			if mf == nil {
				mf = eq
			}
			if got := exp.Must(strconv.Atoi(v)); !mf(got, want[k]) {
				match = false
				break
			}
		}

		if match {
			fmt.Println(num)
		}
	}
}
