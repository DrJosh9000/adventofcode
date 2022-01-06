package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/12.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	var pos [][3]int
	for i := 0; i < 4; i++ {
		var p [3]int
		if _, err := fmt.Fscanf(f, "<x=%d, y=%d, z=%d>\n", &p[0], &p[1], &p[2]); err != nil {
			log.Fatalf("Couldn't scan position: %v", err)
		}
		pos = append(pos, p)
	}

	// all the axes are independent, so simulate them independently
	type state [8]int
	next := func(st state) state {
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				d := sign(st[i] - st[j])
				st[i+4] -= d
				st[j+4] += d
			}
		}
		for i := 0; i < 4; i++ {
			st[i] += st[i+4]
		}
		return st
	}

	period := func(initial state) int {
		p := 1
		for st := next(initial); st != initial; st = next(st) {
			p++
		}
		return p
	}

	x := period(state{pos[0][0], pos[1][0], pos[2][0], pos[3][0]})
	y := period(state{pos[0][1], pos[1][1], pos[2][1], pos[3][1]})
	z := period(state{pos[0][2], pos[1][2], pos[2][2], pos[3][2]})
	t := x * (y / gcd(x, y))
	fmt.Println(t * (z / gcd(t, z)))
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
