package main

import (
	"errors"
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 24, part b

func main() {
	input := exp.MustReadByteGrid("inputs/24.txt")

	var start, end image.Point
	for x, c := range input[0] {
		if c == '.' {
			start.X = x
			break
		}
	}
	end.Y = len(input) - 1
	for x, c := range input[len(input)-1] {
		if c == '.' {
			end.X = x
		}
	}

	storm := make([]grid.Sparse[[]image.Point], 1)
	storm[0] = grid.MapToSparse(input, func(b byte) ([]image.Point, bool) {
		if b == '.' || b == '#' {
			return nil, false
		}
		return []image.Point{algo.CGVL[rune(b)]}, true
	})

	valley := input.Bounds().Inset(1)
	dx, dy := valley.Dx(), valley.Dy()
	evolve := func() {
		st2 := make(grid.Sparse[[]image.Point], len(storm[0]))
		for p, ds := range storm[len(storm)-1] {
			for _, d := range ds {
				q := p.Add(d)
				if !q.In(valley) {
					q.X = (q.X-1+dx)%dx + 1
					q.Y = (q.Y-1+dy)%dy + 1
				}
				st2[q] = append(st2[q], d)
			}
		}

		storm = append(storm, st2)
	}

	type state struct {
		p        image.Point
		t        int
		wp1, wp2 bool
	}
	steps := append(algo.Neigh4, image.Point{})
	startToEnd := algo.L1(start.Sub(end))

	algo.AStar(state{p: start}, func(s state) int {
		switch {
		case s.wp1 && s.wp2:
			return algo.L1(s.p.Sub(end))
		case s.wp1:
			return startToEnd + algo.L1(s.p.Sub(start))
		default:
			return 2*startToEnd + algo.L1(s.p.Sub(end))
		}
	}, func(s state, t int) (map[state]int, error) {
		switch {
		case s.p == end && s.wp1 && s.wp2:
			fmt.Println(t)
			return nil, errors.New("all done")
		case s.p == start && s.wp1:
			s.wp2 = true
		case s.p == end:
			s.wp1 = true
		}

		t++
		for t >= len(storm) {
			evolve()
		}
		next := make(map[state]int)
		for _, d := range steps {
			q := s.p.Add(d)
			if !q.In(valley) && q != start && q != end {
				continue
			}
			if len(storm[t][q]) > 0 {
				continue
			}
			s2 := s
			s2.p = q
			s2.t = t
			next[s2] = 1
		}
		return next, nil
	})
}
