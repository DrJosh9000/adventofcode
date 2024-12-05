package main

import (
	"fmt"
	"image"
	"slices"
	"sync/atomic"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
	"drjosh.dev/exp/para"
)

// Advent of Code 2023
// Day 22, part b

const inputPath = "2023/inputs/22.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	type brick struct {
		start, end algo.Vec3[int]
		above      algo.Set[*brick] // set of bricks above this brick
		below      algo.Set[*brick] // set of bricks below this brick
		supBy      algo.Set[*brick] // subset of below: bricks supporting this brick
		suping     algo.Set[*brick] // subset of above: bricks this brick supports
	}
	bricks := algo.Map(lines, func(line string) (b *brick) {
		b = new(brick)
		exp.MustSscanf(line, "%d,%d,%d~%d,%d,%d", &b.start[0], &b.start[1], &b.start[2], &b.end[0], &b.end[1], &b.end[2])
		// if b.end[2] < b.start[2] {
		// 	fmt.Println("end below start?")
		// }
		return b
	})

	stacks := make(map[image.Point]algo.Set[*brick])
	for _, b := range bricks {
		v := b.start
		s := b.end.Sub(b.start)
		if l1 := s.L1(); l1 != 0 {
			s = s.Div(l1)
		}
		for {
			p := image.Pt(v[0], v[1])
			stacks[p] = stacks[p].Insert(b)
			if v == b.end {
				break
			}
			v = v.Add(s)
		}
	}

	for _, sst := range stacks {
		if len(sst) == 0 {
			continue
		}
		stack := sst.ToSlice()
		slices.SortFunc(stack, func(a, b *brick) int {
			return a.start[2] - b.start[2]
		})
		for i, above := range stack[1:] {
			below := stack[i]
			above.below = above.below.Insert(below)
			below.above = below.above.Insert(above)
		}
	}

	var queue []*brick
	tally := make(map[*brick]int)
	for _, b := range bricks {
		tally[b] = len(b.below)
		if len(b.below) == 0 {
			queue = append(queue, b)
		}
	}

	for len(queue) > 0 {
		b := queue[0]
		queue = queue[1:]

		z := 1
		for below := range b.below {
			t := below.end[2] + 1
			switch {
			case t > z:
				clear(b.supBy)
				b.supBy = b.supBy.Insert(below)
				z = t
			case t == z:
				b.supBy = b.supBy.Insert(below)
			}
		}
		for sup := range b.supBy {
			sup.suping = sup.suping.Insert(b)
		}

		dz := b.end[2] - b.start[2]
		b.start[2] = z
		b.end[2] = z + dz

		for above := range b.above {
			tally[above]--
			if tally[above] == 0 {
				queue = append(queue, above)
			}
		}
	}

	var count atomic.Int64
	para.Do(bricks, func(b *brick) {
		gone := make(algo.Set[*brick])
		queue := []*brick{b}
		for len(queue) > 0 {
			b := queue[0]
			queue = queue[1:]
			gone.Insert(b)

			for a := range b.suping {
				tally := len(a.supBy)
				for sb := range a.supBy {
					if gone.Contains(sb) {
						tally--
					}
				}
				if tally == 0 {
					queue = append(queue, a)
				}
			}
		}
		count.Add(int64(len(gone) - 1))
	})

	fmt.Println(count.Load())
}
