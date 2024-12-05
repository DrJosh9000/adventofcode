package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2023
// Day 11, part a

const inputPath = "2023/inputs/11.txt"

func main() {
	g := exp.MustReadByteGrid(inputPath)
	h, w := g.Size()
	var galaxies []image.Point
	emptyRows := make([]bool, h)
	emptyCols := make([]bool, w)
	for y := range g {
		emptyRows[y] = true
	}
	for x := range g[0] {
		emptyCols[x] = true
	}

	for y, row := range g {
		for x, c := range row {
			if c == '#' {
				galaxies = append(galaxies, image.Pt(x, y))
				emptyRows[y] = false
				emptyCols[x] = false
			}
		}
	}

	er, ec := make([]int, h), make([]int, w)
	acc := 0
	for y, b := range emptyRows {
		if b {
			acc++
		}
		er[y] = acc
	}
	acc = 0
	for x, b := range emptyCols {
		if b {
			acc++
		}
		ec[x] = acc
	}

	sum := 0
	for i, p := range galaxies[:len(galaxies)-1] {
		for _, q := range galaxies[i+1:] {
			sum += algo.L1(p.Sub(q)) + algo.Abs(er[p.Y]-er[q.Y]) + algo.Abs(ec[p.X]-ec[q.X])
		}
	}

	fmt.Println(sum)
}
