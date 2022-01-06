package main

import (
	"bufio"
	"image"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("inputs/10.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	ims := image.Pt(31, 20)

	var asts []image.Point
	sc := bufio.NewScanner(f)
	j := 0
	for sc.Scan() {
		for i, c := range sc.Text() {
			p := image.Pt(i, j).Sub(ims)
			if c != '#' || p == image.Pt(0, 0) {
				continue
			}
			asts = append(asts, p)
		}
		j++
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	sort.Slice(asts, func(i, j int) bool {
		l, r := asts[i].X*asts[j].Y, asts[j].X*asts[i].Y
		if l == r {
			return norm(asts[i]) < norm(asts[j])
		}
		return l > r
	})

	// TODO
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

func norm(p image.Point) int { return abs(p.X) + abs(p.Y) }
