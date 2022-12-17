package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 17, part b

var rocks = []grid.Dense[bool]{
	{
		{true, true, true, true},
	},
	{
		{false, true, false},
		{true, true, true},
		{false, true, false},
	},
	{
		{true, true, true},
		{false, false, true},
		{false, false, true},
	},
	{
		{true},
		{true},
		{true},
		{true},
	},
	{
		{true, true},
		{true, true},
	},
}

const width = 7

func collide(arena grid.Sparse[bool], piece grid.Dense[bool], pos image.Point) bool {
	pb := piece.Bounds()
	if pos.X < 0 {
		return true
	}
	if pos.X+pb.Dx() > width {
		return true
	}
	if pos.Y < 0 {
		return true
	}

	var q image.Point
	for q.Y = pb.Min.Y; q.Y < pb.Max.Y; q.Y++ {
		for q.X = pb.Min.X; q.X < pb.Max.X; q.X++ {
			if !piece[q.Y][q.X] {
				continue
			}
			if arena[q.Add(pos)] {
				return true
			}
		}
	}
	return false
}

func stencil(arena grid.Sparse[bool], piece grid.Dense[bool], pos image.Point) {
	pb := piece.Bounds()

	var q image.Point
	for q.Y = pb.Min.Y; q.Y < pb.Max.Y; q.Y++ {
		for q.X = pb.Min.X; q.X < pb.Max.X; q.X++ {
			if piece[q.Y][q.X] {
				arena[q.Add(pos)] = true
			}
		}
	}
}

var bump = map[byte]image.Point{
	'<': image.Pt(-1, 0),
	'>': image.Pt(1, 0),
}

func main() {
	push := bytes.TrimSpace(exp.Must(os.ReadFile("inputs/17.txt")))
	arena := make(grid.Sparse[bool])

	N := 1000000000000

	trial := len(rocks) * len(push)
	heights := make([]int, 10*trial)
	height := 0
	jet := 0
	for r := range heights {
		piece := rocks[r%len(rocks)]
		p := image.Pt(2, height+3)
		for {
			b := bump[push[jet]]
			if b.X == 0 {
				log.Fatalf("Unknown character in input %c at %d", push[jet], jet)
			}
			q := p.Add(b)
			if !collide(arena, piece, q) {
				p = q
			}
			jet++
			jet %= len(push)

			q = p.Add(image.Pt(0, -1))
			if collide(arena, piece, q) {
				stencil(arena, piece, p)
				if h := p.Y + piece.Bounds().Dy(); h > height {
					height = h
				}
				heights[r] = height
				break
			} else {
				p = q
			}
		}
	}

	nh := len(heights)
perLoop:
	for per := 1; per < nh; per++ {
		d := heights[nh-1] - heights[nh-1-per]
		for i := range heights {
			j := nh - i - 1
			if heights[j]-heights[j-per] != d {
				continue perLoop
			}
			if i > 9*trial {
				fmt.Printf("correlated per=%d, d=%d\n", per, d)
				base := (nh/per - 1) * per
				N -= base
				// why -1?
				fmt.Println(d*(N/per) + heights[N%per+base] - 1)
				return
			}
		}
	}

}
