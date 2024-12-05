package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"

	"drjosh.dev/exp"
	"drjosh.dev/exp/grid"
)

// Advent of Code 2022
// Day 22, part a

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

	board.Map(func(b byte) byte {
		if b == 0 {
			return ' '
		}
		return b
	})

	bounds := board.Bounds()
	var p image.Point
	for x, c := range board[0] {
		if c == '.' {
			p.X = x
			break
		}
	}

	// fmt.Println(len(board))
	// fmt.Println(len(instr))
	// fmt.Println(p)

	tokenRE := regexp.MustCompile(`L|R|\d+`)
	tokens := tokenRE.FindAllString(instr, -1)
	// fmt.Println(tokens)

	d := image.Pt(1, 0)
	for _, token := range tokens {
		switch token {
		case "L":
			d = image.Pt(d.Y, -d.X)
			// fmt.Println("facing", d)
		case "R":
			d = image.Pt(-d.Y, d.X)
			// fmt.Println("facing", d)
		default:
			n := exp.Must(strconv.Atoi(token))
		moveLoop:
			for i := 0; i < n; i++ {
				q := p
			stepLoop:
				for {
					q = q.Add(d)
					if !q.In(bounds) {
						q.X = (q.X + bounds.Dx()) % bounds.Dx()
						q.Y = (q.Y + bounds.Dy()) % bounds.Dy()
					}
					switch board[q.Y][q.X] {
					case '.':
						p = q
						break stepLoop
					case '#':
						// p is unchanged
						break moveLoop
					case ' ':
						// keep going
					}
				}
			}
			// fmt.Println("moved to", p)
		}
	}

	facing := map[image.Point]int{
		{1, 0}:  0,
		{0, 1}:  1,
		{-1, 0}: 2,
		{0, -1}: 3,
	}
	// fmt.Println(p, d)
	fmt.Println(1000*(p.Y+1) + 4*(p.X+1) + facing[d])
}
