package main

import (
	"bytes"
	"fmt"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2016
// Day 21, part b

func main() {
	// to invert "rotate based" we need to know where the letter was *before*
	// the rotation...
	// there are 8 letters, so letter x can be in index {0..7}.
	// if it began index {0,1,2,3}, the string was rrot by {1,2,3,4},
	// ending up in index {1,3,5,7}.
	// if it began index {4,5,6,7}, the string was rrot {6,7,8,9},
	// ending up in index {10,12,14,16}%8 i.e. {2,4,6,0}.
	// hence:
	// left rotation amount for index
	inverse := []int{
		0: 1, // 9
		1: 1,
		2: 6,
		3: 2,
		4: 7,
		5: 3,
		6: 0, // 8
		7: 4,
	}

	pw := []byte("fbgdceah")
	lines := exp.MustReadLines("inputs/21.txt")
	algo.Reverse(lines)
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "swap position"):
			// swap is self-inverse
			var x, y int
			exp.Must(fmt.Sscanf(line, "swap position %d with position %d", &x, &y))
			pw[x], pw[y] = pw[y], pw[x]
		case strings.HasPrefix(line, "swap letter"):
			// swap is self-inverse
			var x, y byte
			exp.Must(fmt.Sscanf(line, "swap letter %c with letter %c", &x, &y))
			i, j := bytes.IndexByte(pw, x), bytes.IndexByte(pw, y)
			pw[i], pw[j] = pw[j], pw[i]
		case strings.HasPrefix(line, "rotate left"):
			// the inverse of rotate left is rotate right
			var x int
			exp.Must(fmt.Sscanf(line, "rotate left %d step", &x))
			pw2 := append([]byte(nil), pw...)
			for i := range pw {
				pw2[(i+x)%len(pw)] = pw[i]
			}
			pw = pw2
		case strings.HasPrefix(line, "rotate right"):
			// the inverse of rotate right is rotate left
			var x int
			exp.Must(fmt.Sscanf(line, "rotate right %d step", &x))
			pw2 := append([]byte(nil), pw...)
			for i := range pw {
				pw2[i] = pw[(i+x)%len(pw)]
			}
			pw = pw2
		case strings.HasPrefix(line, "rotate based"):
			// left rotation. see discussion above.
			var x byte
			exp.Must(fmt.Sscanf(line, "rotate based on position of letter %c", &x))
			y := inverse[bytes.IndexByte(pw, x)]
			pw2 := append([]byte(nil), pw...)
			for i := range pw {
				pw2[i] = pw[(i+y)%len(pw)]
			}
			pw = pw2
		case strings.HasPrefix(line, "reverse"):
			// reverse is self-inverse
			var x, y int
			exp.Must(fmt.Sscanf(line, "reverse positions %d through %d", &x, &y))
			for x < y {
				pw[x], pw[y] = pw[y], pw[x]
				x++
				y--
			}
		case strings.HasPrefix(line, "move"):
			// the inverse of move x to y is move y to x
			var x, y int
			exp.Must(fmt.Sscanf(line, "move position %d to position %d", &y, &x))
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
