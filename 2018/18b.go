package main

import (
	"fmt"
	"image"
	"strings"

	"github.com/DrJosh9000/exp"
)

var neighs = []image.Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, /* me */  {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func main() {
	state := make(map[image.Point]rune)
	y := 0
	exp.MustForEachLineIn("inputs/18.txt", func(line string) {
		for x, c := range line {
			state[image.Pt(x, y)] = c
		}
		y++
	})
	
	// ...looks periodic, after a bit...
	score := func() int {
		trees, lys := 0, 0
		for _, c := range state {
			switch c {
			case '|':
				trees++
			case '#':
				lys++
			}
		}
		return trees * lys
	}
	
	stringise := func() string {
		var sb strings.Builder
		for y := 0; y < 50; y++ {
			for x := 0; x < 50; x++ {
				sb.WriteRune(state[image.Pt(x, y)])
			}
		}
		return sb.String()
	}
	
	var scores []int
	hist := make(map[string]int)
	const target = 1_000_000_000
	
	for m := 0; m < 1000; m++ {
		str := stringise()
		if x, seen := hist[str]; seen {
			period := m - x
			fmt.Println(scores[(target - x) % period + x])
			return
		}
		hist[str] = m
		scores = append(scores, score())
		
		s2 := make(map[image.Point]rune, len(state))
		for p, c := range state {
			s2[p] = c
			switch c {
			case '.':
				// becomes | if 3 or more adjacent areas are |
				tc := 0
				for _, d := range neighs {
					if state[p.Add(d)] == '|' {
						tc++
					}
				}
				if tc >= 3 {
					s2[p] = '|'
				}
			case '|':
				// becomes | if 3 or more adjacent areas are |
				tc := 0
				for _, d := range neighs {
					if state[p.Add(d)] == '#' {
						tc++
					}
				}
				if tc >= 3 {
					s2[p] = '#'
				}
			case '#':
				// becomes . if not next to a | and a #
				tree, ly := false, false
				for _, d := range neighs {
					if state[p.Add(d)] == '|' {
						tree = true
					}
					if state[p.Add(d)] == '#' {
						ly = true
					}
				}
				if !(tree && ly) {
					s2[p] = '.'
				}
			}
		}
		state = s2
	}
	
}