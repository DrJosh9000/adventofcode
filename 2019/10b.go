package main

import (
	"bufio"
	"fmt"
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

	count := 0
	for {
		sort.Slice(asts, func(i, j int) bool {
			qi, qj := quad(asts[i]), quad(asts[j])
			if qi == qj {
				sa := asts[i].X*asts[j].Y - asts[j].X*asts[i].Y
				if sa == 0 {
					return norm(asts[i]) < norm(asts[j])
				}
				return sa > 0
			}
			return qi < qj
		})

		var newasts []image.Point
		var slope image.Point
		for _, p := range asts {
			q := p.Div(gcd(abs(p.X), abs(p.Y)))
			if slope == q {
				newasts = append(newasts, p)
				continue
			}
			slope = q
			count++
			if count == 200 {
				fmt.Println(p.Add(ims))
				return
			}
		}
		asts = newasts
	}
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

func quad(p image.Point) int {
	if p.X >= 0 && p.Y < 0 {
		return 0
	}
	if p.Y >= 0 && p.X > 0 {
		return 1
	}
	if p.Y > 0 && p.X <= 0 {
		return 2
	}
	return 3
}
