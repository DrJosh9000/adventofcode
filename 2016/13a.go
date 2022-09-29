package main

import (
	"errors"
	"fmt"
	"image"
	"math/bits"
	"os"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2016
// Day 13, part a

func main() {
	fav := uint(exp.Must(strconv.Atoi(os.Args[1])))

	start := image.Pt(1, 1)
	goal := image.Pt(31, 39)
	algo.FloodFill(start, func(p image.Point, d int) ([]image.Point, error) {
		if p == goal {
			fmt.Println(d)
			return nil, errors.New("all done")
		}
		var next []image.Point
		for _, s := range algo.Neigh4 {
			q := p.Add(s)
			if q.X < 0 || q.Y < 0 {
				continue
			}
			x, y := uint(q.X), uint(q.Y)
			if bits.OnesCount(x*x+3*x+2*x*y+y+y*y+fav)%2 == 1 {
				continue
			}
			next = append(next, q)
		}
		return next, nil
	})
}
