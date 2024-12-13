package main

import (
	_ "embed"
	"fmt"
	"image"
	"math"
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
			// fmt.Printf("a, b, prize = %v, %v, %v\n", a, b, prize)

			mint := math.MaxInt
			for bp := range 101 {
				for ap := range 101 {
					if a.Mul(ap).Add(b.Mul(bp)) == prize {
						mint = min(mint, 3*ap+bp)
					}
				}
			}
			if mint != math.MaxInt {
				tokens += mint
			}

		default:
			continue
		}

	}
	fmt.Println(tokens)
}
