package main

import (
	"fmt"
	"os"
	"strings"
)

// Advent of Code 2015
// Day 11

const alphabet = "abcdefghjkmnpqrstuvwxyz"
const N = len(alphabet)

func printPW(x int) {
	pw := make([]byte, 8)
	for i := range pw {
		pw[7-i] = alphabet[x%N]
		x /= N
	}
	fmt.Printf("%s\n", pw)
}

func search(x int) int {
	for {
		x++

		// look for straights
		y := x
		str8 := false
		for i := 0; i < 6 && !str8; i++ {
			if alphabet[y%N] != alphabet[(y/N)%N]+1 {
				y /= N
				continue
			}
			y /= N
			if alphabet[y%N] != alphabet[(y/N)%N]+1 {
				continue
			}
			str8 = true
		}
		if !str8 {
			continue
		}

		// look for pairs
		y = x
		p1 := -1
		for i := 0; i < 7 && p1 == -1; i++ {
			if y%N != (y/N)%N {
				y /= N
				continue
			}
			p1 = y % N
		}

		if p1 == -1 {
			// no pair
			continue
		}

		y = x
		p2 := -1
		for i := 0; i < 7 && p2 == -1; i++ {
			if y%N == p1 {
				y /= N
				continue
			}
			if y%N != (y/N)%N {
				y /= N
				continue
			}
			p2 = y % N
		}
		if p2 == -1 {
			// no other pair
			continue
		}

		return x
	}
}

func main() {
	x := 0
	for _, c := range os.Args[1] {
		x *= N
		x += strings.IndexRune(alphabet, c)
	}

	x = search(x)
	printPW(x)
	x = search(x)
	printPW(x)
}
