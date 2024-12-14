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
	robots := algo.Map(exp.NonEmpty(strings.Split(input, "\n")), func(line string) *robot {
		r := new(robot)
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.p.X, &r.p.Y, &r.v.X, &r.v.Y)
		return r
	})

	bounds := image.Rect(0, 0, 101, 103)
	centre := image.Rect(35, 31, 66, 72)
	it := 0
	for {
		it++
		count := 0
		for _, r := range robots {
			r.p = r.p.Add(r.v).Mod(bounds)
			if r.p.In(centre) {
				count++
			}
		}
		if count > 200 {
			fmt.Println(it)
			return
		}
	}

}
