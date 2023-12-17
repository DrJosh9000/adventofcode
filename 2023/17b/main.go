package main

import (
	"errors"
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 17, part b

const inputPath = "2023/inputs/17.txt"

func main() {
	g := exp.MustReadByteGrid(inputPath)
	bounds := g.Bounds()
	goal := bounds.Max.Add(image.Pt(-1, -1))
	type state struct {
		p, d image.Point
		n    int
	}
	done := errors.New("done")
	algo.AStar(
		state{},
		func(s state) int {
			return algo.L1(goal.Sub(s.p))
		},
		func(s state, d int) (map[state]int, error) {
			if s.p == goal {
				fmt.Println(d)
				return nil, done
			}
			var ds []image.Point
			switch s.n {
			case 0:
				ds = algo.Neigh4
			case 1, 2, 3:
				ds = []image.Point{s.d}
			case 4, 5, 6, 7, 8, 9:
				ds = []image.Point{s.d, {s.d.Y, -s.d.X}, {-s.d.Y, s.d.X}}
			case 10:
				ds = []image.Point{{s.d.Y, -s.d.X}, {-s.d.Y, s.d.X}}
			}
			next := make(map[state]int)
			for _, d := range ds {
				q := s.p.Add(d)
				if !q.In(bounds) {
					continue
				}
				ns := s
				ns.p = ns.p.Add(d)
				if ns.d == d {
					ns.n++
				} else {
					ns.n = 1
				}
				ns.d = d
				next[ns] = int(g.At(ns.p) - '0')
			}
			return next, nil
		},
	)
}
