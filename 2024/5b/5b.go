package main

import (
	_ "embed"
	"fmt"
	"image"
	"slices"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

//go:embed inputs/5.txt
var input string

func main() {
	pairs := make(algo.Set[image.Point])
	seqs := false
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			seqs = true
			continue
		}
		if seqs {
			seq := algo.Map(strings.Split(line, ","), exp.MustAtoi)
			good := true
		seqLoop:
			for a, b := range algo.AllPairs(seq) {
				if pairs.Contains(image.Pt(b, a)) {
					good = false
					break seqLoop
				}
			}
			if !good {
				slices.SortStableFunc(seq, func(a, b int) int {
					if pairs.Contains(image.Pt(a, b)) {
						return -1
					}
					if pairs.Contains(image.Pt(b, a)) {
						return 1
					}
					return 0
				})

				sum += seq[len(seq)/2]
			}

		} else {
			bef, aft, ok := strings.Cut(line, "|")
			if !ok {
				panic(line)
			}
			pairs.Insert(image.Pt(exp.MustAtoi(bef), exp.MustAtoi(aft)))
		}
	}
	fmt.Println(sum)
}
