package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

//go:embed inputs/14.txt
var input string

type robot struct {
	p, v image.Point
}

func main() {
	robots := algo.Map(exp.NonEmpty(strings.Split(input, "\n")), func(line string) (r robot) {
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.p.X, &r.p.Y, &r.v.X, &r.v.Y)
		return r
	})

	bounds := image.Rect(0, 0, 101, 103)
	for range 100 {
		for i := range robots {
			r := &robots[i]
			r.p = r.p.Add(r.v).Mod(bounds)
		}
	}

	qs := [4]image.Rectangle{
		image.Rect(0, 0, 50, 51),
		image.Rect(51, 0, 101, 51),
		image.Rect(0, 52, 50, 103),
		image.Rect(51, 52, 101, 103),
	}
	s := make([]int, 4)
	for _, r := range robots {
		for i, q := range qs {
			if r.p.In(q) {
				s[i]++
			}
		}
	}
	fmt.Println(algo.Prod(s))
}
