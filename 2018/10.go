package main

import (
	"fmt"
	"image"
	"log"
	"math"

	"drjosh.dev/exp"
)

type star struct {
	pos, vel image.Point
}

func main() {
	var stars []star
	exp.MustForEachLineIn("inputs/10.txt", func(line string) {
		var s star
		if _, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &s.pos.X, &s.pos.Y, &s.vel.X, &s.vel.Y); err != nil {
			log.Fatalf("Couldn't parse line %q: %v", line, err)
		}
		stars = append(stars, s)
	})
	
	h2 := func(t int) int {
		minh, maxh := math.MaxInt, math.MinInt
		for _, s := range stars {
			h := s.pos.Y + t * s.vel.Y
			if h < minh {
				minh = h
			}
			if h > maxh {
				maxh = h
			}
		}
		h := maxh - minh
		return h * h
	}

	a0, a1, a2 := h2(0), h2(1), h2(2)

	// h2(t) is parabola-ish and entirely above the t axis:
	//   h2(t) ~= at^2 + bt + c
	// so
	//   h2(0) ~= c
	//   h2(1) ~= a + b + c
	//   h2(2) ~= 4a + 2b + c
	// hence
	//   h2(2) - 2h2(1) + c ~= (4a + 2b + c) - (2a + 2b + 2c) + c = 2a
	adub := (a2 - 2*a1 + a0)
	// and 
	//   b ~= h2(1) - a - c
	b := a1 - adub/2 - a0
	// Turning point is t = -b / 2a
	t := -b / adub
	
	fmt.Println(t)
	
	bounds := image.Rectangle{
		Min: image.Point{math.MaxInt, math.MaxInt}, 
		Max: image.Point{math.MinInt, math.MinInt},
	}
	for _, s := range stars {
		p := s.pos.Add(s.vel.Mul(t))
		if p.X < bounds.Min.X {
			bounds.Min.X = p.X
		}
		if p.X > bounds.Max.X {
			bounds.Max.X = p.X
		}
		if p.Y < bounds.Min.Y {
			bounds.Min.Y = p.Y
		}
		if p.Y > bounds.Max.Y {
			bounds.Max.Y = p.Y
		}
	}
	m := make([][]byte, bounds.Dy()+1)
	for i := range m {
		m[i] = make([]byte, bounds.Dx()+1)
		for j := range m[i] {
			m[i][j] = ' '
		}
	}
	for _, s := range stars {
		p := s.pos.Add(s.vel.Mul(t)).Sub(bounds.Min)
		m[p.Y][p.X] = '#'
	}
	for i := range m {
		fmt.Println(string(m[i]))
	}
}