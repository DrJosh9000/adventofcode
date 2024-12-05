package main

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"drjosh.dev/exp/algo"
)

func hash(input []byte) []byte {
	input = append(input, 17, 31, 73, 47, 23)

	circle := make([]byte, 256)
	for i := range circle {
		circle[i] = byte(i)
	}
	var pos, skip byte

	for round := 0; round < 64; round++ {
		for _, l := range input {
			for i := byte(0); i < l/2; i++ {
				j, k := pos+i, pos+l-i-1
				circle[j], circle[k] = circle[k], circle[j]
			}
			pos += l + skip
			skip++
		}
	}

	dense := make([]byte, 16)
	for i := 0; i < 16; i++ {
		dense[i] = algo.Foldl(circle[i*16:(i+1)*16], func(x, y byte) byte {
			return x ^ y
		})
	}
	return dense
}

func main() {
	input := os.Args[1]

	d := make(algo.DisjointSets[image.Point])
	for y := 0; y < 128; y++ {
		for x, b := range hash([]byte(input + "-" + strconv.Itoa(y))) {
			for z := 0; z < 8; z++ {
				if b&(1<<z) != 0 {
					p := image.Pt(8*x+(7-z), y)
					d[p] = p
				}
			}
		}
	}
	neigh := []image.Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for p := range d {
		for _, n := range neigh {
			q := p.Add(n)
			if _, adj := d[q]; adj {
				d.Union(p, q)
			}
		}
	}

	fmt.Println(len(d.Reps()))
}
