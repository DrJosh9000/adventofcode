package main

import (
	_ "embed"
	"fmt"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

//go:embed inputs/7.txt
var input string

func concat(a, b int) int {
	for c := b; c > 0; c /= 10 {
		a *= 10
	}
	return a + b
}

func main() {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		anss, valss, ok := strings.Cut(line, ": ")
		if !ok {
			continue
		}
		ans := exp.MustAtoi(anss)
		vals := algo.Map(strings.Fields(valss), exp.MustAtoi)
		var rec func(idx, acc int) bool
		rec = func(idx, acc int) bool {
			if idx == len(vals) {
				return acc == ans
			}
			return rec(idx+1, acc+vals[idx]) || rec(idx+1, acc*vals[idx]) || rec(idx+1, concat(acc, vals[idx]))
		}
		if rec(1, vals[0]) {
			sum += ans
		}
	}
	fmt.Println(sum)
}
