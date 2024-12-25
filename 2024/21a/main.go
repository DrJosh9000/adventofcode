package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"drjosh.dev/exp"
)

// Advent of Code 2024
// Day 21, part a

const inputPath = "2024/inputs/21.txt"

func main() {
	lines := exp.MustReadLines(inputPath)
	sum := 0
	for _, line := range lines {
		code := int(exp.Must(strconv.ParseInt(line[:3], 10, 64)))
		p := numpad['A']
		var sb strings.Builder
		for _, c := range line {
			q := numpad[c]
			switch {
			case p.X == q.X:
				sb.WriteString(move(p, q))
			case p.Y == q.Y:
				sb.WriteString(move(p, q))
			case p.Y == 3 && q.X == 0:
				sb.WriteString(move(p, q))
			case p.X == 0 && q.Y == 3:
				sb.WriteString(move(p, q))
			default:
				m1, m2 := move(p, q), move2(p, q)
				e1 := expand(expand(m1))
				e2 := expand(expand(m2))
				if len(e1) < len(e2) {
					sb.WriteString(m1)
				} else {
					sb.WriteString(m2)
				}
			}
			p = q
		}
		sum += code * len(expand(expand(sb.String())))
	}
	fmt.Println(sum)
}

func expand(dirs string) string {
	var sb strings.Builder
	p := dirpad['A']
	for _, c := range dirs {
		q := dirpad[c]
		sb.WriteString(move(p, q))
		p = q
	}
	return sb.String()
}

func move(p, q image.Point) string {
	var sb strings.Builder
	for p.X < q.X {
		sb.WriteRune('>')
		p.X++
	}
	for p.Y < q.Y {
		sb.WriteRune('v')
		p.Y++
	}
	for p.Y > q.Y {
		sb.WriteRune('^')
		p.Y--
	}
	for p.X > q.X {
		sb.WriteRune('<')
		p.X--
	}
	sb.WriteRune('A')
	return sb.String()
}

func move2(p, q image.Point) string {
	var sb strings.Builder
	for p.X > q.X {
		sb.WriteRune('<')
		p.X--
	}
	for p.Y > q.Y {
		sb.WriteRune('^')
		p.Y--
	}
	for p.Y < q.Y {
		sb.WriteRune('v')
		p.Y++
	}
	for p.X < q.X {
		sb.WriteRune('>')
		p.X++
	}
	sb.WriteRune('A')
	return sb.String()
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

/*
func reduce(dirs string) string {
	var sb strings.Builder
	p := dirpad['A']
	for _, c := range dirs {
		if c == 'A' {
			sb.WriteRune(dirpadG.At(p))
		} else {
			p = p.Add(algo.CGVL[c])
		}
	}
	return sb.String()
}

var numpadG = grid.Dense[rune]{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{0, '0', 'A'},
}

var dirpadG = grid.Dense[rune]{
	{0, '^', 'A'},
	{'<', 'v', '>'},
}
*/
