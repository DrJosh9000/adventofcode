package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/grid"
)

// Advent of Code 2022
// Day 22, part b

func main() {
	input := exp.MustReadLines("inputs/22.txt")
	var board grid.Dense[byte]
	var instr string
	for i, line := range input {
		if line == "" {
			board = grid.BytesFromStrings(input[:i])
			instr = input[i+1]
			break
		}
	}

	type posdir struct{ p, d image.Point }

	pd := posdir{d: image.Pt(1, 0)}
	for x, c := range board[0] {
		if c == '.' {
			pd.p.X = x
			break
		}
	}

	const edgelen = 50

	next := make(map[posdir]posdir)
	stitch := func(line1 posdir, outdir image.Point, line2 posdir, indir image.Point) {
		p, q := line1.p, line2.p
		dp, dq := line1.d, line2.d
		for i := 0; i < edgelen; i++ {
			next[posdir{p, outdir}] = posdir{q, indir}
			p = p.Add(dp)
			q = q.Add(dq)
		}
		p, q = line2.p, line1.p
		dp, dq = line2.d, line1.d
		outdir, indir = indir.Mul(-1), outdir.Mul(-1)
		for i := 0; i < edgelen; i++ {
			next[posdir{p, outdir}] = posdir{q, indir}
			p = p.Add(dp)
			q = q.Add(dq)
		}
	}

	// I should figure out a way to automate this...

	mkpd := func(px, py, dx, dy int) posdir {
		return posdir{p: image.Point{px, py}, d: image.Point{dx, dy}}
	}
	stitch(mkpd(50, 0, 1, 0), image.Pt(0, -1), mkpd(0, 150, 0, 1), image.Pt(1, 0))
	stitch(mkpd(100, 0, 1, 0), image.Pt(0, -1), mkpd(0, 199, 1, 0), image.Pt(0, -1))
	stitch(mkpd(149, 0, 0, 1), image.Pt(1, 0), mkpd(99, 149, 0, -1), image.Pt(-1, 0))
	stitch(mkpd(100, 49, 1, 0), image.Pt(0, 1), mkpd(99, 50, 0, 1), image.Pt(-1, 0))
	stitch(mkpd(50, 149, 1, 0), image.Pt(0, 1), mkpd(49, 150, 0, 1), image.Pt(-1, 0))
	stitch(mkpd(0, 100, 1, 0), image.Pt(0, -1), mkpd(50, 50, 0, 1), image.Pt(1, 0))
	stitch(mkpd(50, 0, 0, 1), image.Pt(-1, 0), mkpd(0, 149, 0, -1), image.Pt(1, 0))

	tokenRE := regexp.MustCompile(`L|R|\d+`)
	tokens := tokenRE.FindAllString(instr, -1)

	for _, token := range tokens {
		switch token {
		case "L":
			pd.d = image.Pt(pd.d.Y, -pd.d.X)
		case "R":
			pd.d = image.Pt(-pd.d.Y, pd.d.X)
		default:
			n := exp.Must(strconv.Atoi(token))
		moveLoop:
			for i := 0; i < n; i++ {
				qd, ok := next[pd]
				if !ok {
					qd = posdir{pd.p.Add(pd.d), pd.d}
				}
				switch board[qd.p.Y][qd.p.X] {
				case '.':
					pd = qd
				case '#':
					break moveLoop
				}
			}
		}
	}

	facing := map[image.Point]int{
		{1, 0}:  0,
		{0, 1}:  1,
		{-1, 0}: 2,
		{0, -1}: 3,
	}
	fmt.Println(1000*(pd.p.Y+1) + 4*(pd.p.X+1) + facing[pd.d])
}
