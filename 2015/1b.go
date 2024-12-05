package main

import (
	"fmt"
	"os"

	"drjosh.dev/exp"
)

// Advent of Code 2015
// Day 1, part b

func main() {
	f := 0
	for i, c := range exp.Must(os.ReadFile("inputs/1.txt")) {
		switch c {
		case '(':
			f++
		case ')':
			f--
		}
		if f == -1 {
			fmt.Println(i + 1)
			return
		}
	}
}
