package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 21, part b

const inputPath = "2024/inputs/21.txt"

func main() {
	for _, a := range []rune{'<', '^', 'v', '>', 'A'} {
		for _, b := range []rune{'<', '^', 'v', '>', 'A'} {
			fragments[fmt.Sprintf("%c%c", a, b)] = frags(move(dirpad[a], dirpad[b], image.Pt(0, 0)))
		}
	}
	// fmt.Println(fragments)

	lines := exp.MustReadLines(inputPath)
	sum := 0
	for _, line := range lines {
		code := int(exp.Must(strconv.ParseInt(line[:3], 10, 64)))
		p := numpad['A']
		var sb strings.Builder
		for _, c := range line {
			q := numpad[c]
			sb.WriteString(move(p, q, image.Pt(0, 3)))
			p = q
		}
		inst := sb.String()
		fm := frags(inst)
		for range 25 {
			fm = expfrags(fm)
		}
		cplx := 0
		for _, n := range fm {
			cplx += n
		}
		sum += code * cplx
	}
	fmt.Println(sum)
}

var fragments = make(map[string]map[string]int)

func frags(s string) map[string]int {
	fm := make(map[string]int)
	s = "A" + s
	for i := range s[1:] {
		fm[s[i:][:2]]++
	}
	return fm
}

func expfrags(m map[string]int) map[string]int {
	m2 := make(map[string]int)
	for f, n := range m {
		for f2, n2 := range fragments[f] {
			m2[f2] += n * n2
		}
	}
	return m2
}

func expand(dirs string) string {
	var sb strings.Builder
	p := dirpad['A']
	for _, c := range dirs {
		q := dirpad[c]
		sb.WriteString(move(p, q, image.Pt(0, 0)))
		p = q
	}
	return sb.String()
}

func move(p, q, g image.Point) string {
	var v, h strings.Builder
	for y := p.Y; y < q.Y; y++ {
		v.WriteRune('v')
	}
	for y := p.Y; y > q.Y; y-- {
		v.WriteRune('^')
	}
	for x := p.X; x < q.X; x++ {
		h.WriteRune('>')
	}
	for x := p.X; x > q.X; x-- {
		h.WriteRune('<')
	}
	if p.X < q.X && image.Pt(p.X, q.Y) != g {
		return v.String() + h.String() + "A"
	}
	if image.Pt(q.X, p.Y) != g {
		return h.String() + v.String() + "A"
	}
	return v.String() + h.String() + "A"
}

var numpad = []image.Point{
	'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
	'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
	'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
	/*        */ '0': {1, 3}, 'A': {2, 3},
}

var dirpad = []image.Point{
	/*        */ '^': {1, 0}, 'A': {2, 0},
	'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
}
