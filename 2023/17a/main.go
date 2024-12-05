package main

import (
	"errors"
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2023
// Day 17, part a

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
			switch {
			case s.n == 0:
				ds = algo.Neigh4
			case s.n == 3:
				ds = []image.Point{{s.d.Y, -s.d.X}, {-s.d.Y, s.d.X}}
			default:
				ds = []image.Point{s.d, {s.d.Y, -s.d.X}, {-s.d.Y, s.d.X}}
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
