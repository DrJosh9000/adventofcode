package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/10.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	asts := make(map[image.Point]struct{})
	sc := bufio.NewScanner(f)
	j := 0
	for sc.Scan() {
		for i, c := range sc.Text() {
			if c != '#' {
				continue
			}
			asts[image.Pt(i, j)] = struct{}{}
		}
		j++
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	best := 0
	var bestp image.Point
	for p := range asts {
		visible := 0
	checkQLoop:
		for q := range asts {
			d := q.Sub(p)
			if g := gcd(abs(d.X), abs(d.Y)); g != 0 {
				d = d.Div(g)
			}
			for t := p.Add(d); t != q; t = t.Add(d) {
				if _, ast := asts[t]; ast {
					continue checkQLoop
				}
			}
			visible++
		}
		if visible > best {
			best = visible
			bestp = p
		}
	}
	fmt.Println(best-1, bestp)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
