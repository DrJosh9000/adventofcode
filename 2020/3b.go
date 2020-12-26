package main

import (
	"bufio"
	"fmt"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.3")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	type sim struct {
		dx, dy, x, y int
		trees int64
	}
	sims := []*sim{
		{dx: 1, dy: 1},
		{dx: 3, dy: 1},
		{dx: 5, dy: 1},
		{dx: 7, dy: 1},
		{dx: 1, dy: 2},
	}
	sc := bufio.NewScanner(f)
	line := 0
	for sc.Scan() {
		if line == 0 {
			// First line is not needed
			line++
			continue
		}
		m := sc.Text()
		for _, sim := range sims {
			if line%sim.dy != 0 {
				continue
			}
			sim.x += sim.dx
			sim.y += sim.dy
			if m[sim.x%len(m)] == '#' {
				sim.trees++
			}
		}
		line++
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	prod := int64(1)
	for _, sim := range sims {
		prod *= sim.trees
	}
	fmt.Println(prod)
}
