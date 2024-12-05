package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2023
// Day 18, part b

const inputPath = "2023/inputs/18.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	type inst struct {
		d image.Point
		n int
	}
	N := len(lines)
	insts := make([]inst, 0, N)
	for _, line := range lines {
		bits := strings.Fields(line)
		c := strings.TrimSuffix(strings.TrimPrefix(bits[2], "(#"), ")")

		n := exp.Must(strconv.ParseInt(c[:5], 16, 64))
		var d image.Point
		switch c[5] {
		case '0':
			d = image.Pt(1, 0)
		case '1':
			d = image.Pt(0, 1)
		case '2':
			d = image.Pt(-1, 0)
		case '3':
			d = image.Pt(0, -1)
		}
		insts = append(insts, inst{
			d: d,
			n: int(n),
		})
	}

	ol, or := corners(insts[N-1].d, insts[0].d)
	var p image.Point
	suml, sumr := 0, 0
	for i, inst := range insts {
		p = p.Add(inst.d.Mul(inst.n))
		l, r := corners(inst.d, insts[(i+1)%N].d)
		pl, pr := p.Add(l), p.Add(r)
		suml += sa2(ol, pl)
		sumr += sa2(or, pr)
		ol, or = pl, pr
	}

	fmt.Println(max(suml, sumr) / 2)
}

func sa2(p, q image.Point) int { return p.X*q.Y - q.X*p.Y }

func corners(d1, d2 image.Point) (l, r image.Point) {
	switch d1 {
	case image.Pt(1, 0): // R
		switch d2 {
		case image.Pt(0, -1): // RU
			return image.Pt(0, 0), image.Pt(1, 1)
		case image.Pt(0, 1): // RD
			return image.Pt(1, 0), image.Pt(0, 1)
		}

	case image.Pt(0, 1): // D
		switch d2 {
		case image.Pt(1, 0): // DR
			return image.Pt(1, 0), image.Pt(0, 1)
		case image.Pt(-1, 0): // DL
			return image.Pt(1, 1), image.Pt(0, 0)
		}

	case image.Pt(-1, 0): // L
		switch d2 {
		case image.Pt(0, -1): // LU
			return image.Pt(0, 1), image.Pt(1, 0)
		case image.Pt(0, 1): // LD
			return image.Pt(1, 1), image.Pt(0, 0)
		}

	case image.Pt(0, -1): // U
		switch d2 {
		case image.Pt(1, 0): // UR
			return image.Pt(0, 0), image.Pt(1, 1)
		case image.Pt(-1, 0): // UL
			return image.Pt(0, 1), image.Pt(1, 0)
		}
	}

	panic(fmt.Sprintf("invalid turn %v %v", d1, d2))
}
