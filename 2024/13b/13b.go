package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	"drjosh.dev/exp"
)

//go:embed inputs/13.txt
var input string

func main() {
	var a, b, prize image.Point
	tokens := 0
	for _, line := range strings.Split(input, "\n") {
		switch {
		case exp.Smatchf(line, "Button A: X+%d, Y+%d", &a.X, &a.Y):
			continue
		case exp.Smatchf(line, "Button B: X+%d, Y+%d", &b.X, &b.Y):
			continue
		case exp.Smatchf(line, "Prize: X=%d, Y=%d", &prize.X, &prize.Y):
			prize.X += 10000000000000
			prize.Y += 10000000000000

			// [ a_x b_x ] [ a_p ] = [ prize_x ]
			// [ a_y b_y ] [ b_p ]   [ prize_y ]
			//
			// [ a_p ] =  1  [  b_y -b_x ] [ prize_x ]
			// [ b_p ]   det [ -a_y  a_x ] [ prize_y ]

			det := a.X*b.Y - b.X*a.Y
			r1 := b.Y*prize.X - b.X*prize.Y
			r2 := a.X*prize.Y - a.Y*prize.X
			if r1%det != 0 || r2%det != 0 {
				continue
			}
			ap, bp := r1/det, r2/det
			if ap < 0 || bp < 0 {
				continue
			}
			tokens += 3*ap + bp

		default:
			continue
		}

	}
	fmt.Println(tokens)
}
