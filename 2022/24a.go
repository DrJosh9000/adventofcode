package main

import (
	"errors"
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2022
// Day 24, part a

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

		// st2d, _ := st2.ToDense()
		// fmt.Println(grid.Map(st2d, func(ps []image.Point) byte {
		// 	if len(ps) == 0 {
		// 		return '.'
		// 	}
		// 	return '0' + byte(len(ps))
		// }))

		storm = append(storm, st2)
	}

	type state struct {
		p image.Point
		t int
	}
	steps := append(algo.Neigh4, image.Point{})

	algo.AStar(state{p: start}, func(s state) int {
		return algo.L1(s.p.Sub(end))
	}, func(s state, t int) (map[state]int, error) {
		//fmt.Println(p, t)
		if s.p == end {
			fmt.Println(t)
			return nil, errors.New("all done")
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
			next[state{p: q, t: t}] = 1
		}
		//fmt.Println(next)
		return next, nil
	})

	// p := end
	// for p != start {
	// 	fmt.Println(p)
	// 	p = prev[p]
	// }
	// fmt.Println(start)
}
