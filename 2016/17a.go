package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"os"
)

// Advent of Code 2016
// Day 17, part a

func main() {
	passcode := append(make([]byte, 0, 32), []byte(os.Args[1])...)

	bounds := image.Rect(0, 0, 4, 4)
	type state struct {
		pos  image.Point
		path string
	}
	q := []state{{}}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		if s.pos == image.Pt(3, 3) {
			fmt.Println(s.path)
			return
		}

		h := md5.Sum(append(passcode, []byte(s.path)...))

		dirs := make(map[string]image.Point)
		if h[0] >= 0xb0 {
			dirs["U"] = image.Pt(0, -1)
		}
		if (h[0] & 0xf) >= 0xb {
			dirs["D"] = image.Pt(0, 1)
		}
		if h[1] >= 0xb0 {
			dirs["L"] = image.Pt(-1, 0)
		}
		if (h[1] & 0xf) >= 0xb {
			dirs["R"] = image.Pt(1, 0)
		}
		for d, δ := range dirs {
			p := s.pos.Add(δ)
			if !p.In(bounds) {
				continue
			}
			q = append(q, state{
				pos:  p,
				path: s.path + d,
			})
		}
	}
}
