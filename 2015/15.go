package main

import (
	"fmt"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 15, part a

type ingredient struct {
	name string

	cap, dur, fla, tex, cal int
}

func (x ingredient) add(y ingredient) ingredient {
	return ingredient{
		cap: x.cap + y.cap,
		dur: x.dur + y.dur,
		fla: x.fla + y.fla,
		tex: x.tex + y.tex,
		cal: x.cal + y.cal,
	}
}

func (x ingredient) mul(y int) ingredient {
	return ingredient{
		cap: y * x.cap,
		dur: y * x.dur,
		fla: y * x.fla,
		tex: y * x.tex,
		cal: y * x.cal,
	}
}

func (x ingredient) norm() int {
	if x.cap <= 0 || x.dur <= 0 || x.fla <= 0 || x.tex <= 0 {
		return 0
	}
	return x.cap * x.dur * x.fla * x.tex
}

func main() {
	var ingreds []ingredient
	for _, line := range exp.MustReadLines("inputs/15.txt") {
		if line == "" {
			continue
		}
		var ing ingredient
		exp.Must(fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &ing.name, &ing.cap, &ing.dur, &ing.fla, &ing.tex, &ing.cal))
		ingreds = append(ingreds, ing)
	}

	N := len(ingreds)
	var maxs, maxscal int

	var search func([]int, ingredient, int)
	search = func(amt []int, sum ingredient, rem int) {
		if len(amt) == N-1 {
			sum = sum.add(ingreds[N-1].mul(rem))
			s := sum.norm()
			if s > maxs {
				maxs = s
			}
			if sum.cal == 500 && s > maxscal {
				maxscal = s
			}
			return
		}
		for a := 0; a <= rem; a++ {
			search(append(amt, a), sum.add(ingreds[len(amt)].mul(a)), rem-a)
		}
	}

	search(make([]int, 0, N), ingredient{}, 100)

	fmt.Println(maxs, maxscal)
}
