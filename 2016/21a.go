package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/DrJosh9000/exp"
)

// Advent of Code 2016
// Day 21, part a

func main() {
	pw := []byte("abcdefgh")
	for _, line := range exp.MustReadLines("inputs/21.txt") {
		switch {
		case strings.HasPrefix(line, "swap position"):
			var x, y int
			exp.Must(fmt.Sscanf(line, "swap position %d with position %d", &x, &y))
			pw[x], pw[y] = pw[y], pw[x]
		case strings.HasPrefix(line, "swap letter"):
			var x, y byte
			exp.Must(fmt.Sscanf(line, "swap letter %c with letter %c", &x, &y))
			i, j := bytes.IndexByte(pw, x), bytes.IndexByte(pw, y)
			pw[i], pw[j] = pw[j], pw[i]
		case strings.HasPrefix(line, "rotate left"):
			var x int
			exp.Must(fmt.Sscanf(line, "rotate left %d step", &x))
			pw2 := append([]byte(nil), pw...)
			for i := range pw {
				pw2[i] = pw[(i+x)%len(pw)]
			}
			pw = pw2
		case strings.HasPrefix(line, "rotate right"):
			var x int
			exp.Must(fmt.Sscanf(line, "rotate right %d step", &x))
			pw2 := append([]byte(nil), pw...)
			for i := range pw {
				pw2[(i+x)%len(pw)] = pw[i]
			}
			pw = pw2
		case strings.HasPrefix(line, "rotate based"):
			var x byte
			exp.Must(fmt.Sscanf(line, "rotate based on position of letter %c", &x))
			y := bytes.IndexByte(pw, x)
			if y >= 4 {
				y++
			}
			y++
			pw2 := append([]byte(nil), pw...)
			for i := range pw {
				pw2[(i+y)%len(pw)] = pw[i]
			}
			pw = pw2
		case strings.HasPrefix(line, "reverse"):
			var x, y int
			exp.Must(fmt.Sscanf(line, "reverse positions %d through %d", &x, &y))
			for x < y {
				pw[x], pw[y] = pw[y], pw[x]
				x++
				y--
			}
		case strings.HasPrefix(line, "move"):
			var x, y int
			exp.Must(fmt.Sscanf(line, "move position %d to position %d", &x, &y))
			c := pw[x]
			for y != x {
				if y < x {
					pw[x] = pw[x-1]
					x--
				} else {
					pw[x] = pw[x+1]
					x++
				}
			}
			pw[y] = c
		}
	}
	fmt.Println(string(pw))
}
