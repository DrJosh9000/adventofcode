package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"log"
	"os"
)

var offsets = []image.Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {0, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func main() {
	f, err := os.Open("inputs/20.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()
	program := sc.Text()
	sc.Scan() // blank line
	var input []string
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	pic := make(map[image.Point]bool)
	for j, row := range input {
		for i, cell := range row {
			pic[image.Pt(i, j)] = cell == '#'
		}
	}

	index := func(p image.Point, infinity bool) int {
		n := 0
		for _, o := range offsets {
			n *= 2
			lit, known := pic[p.Add(o)]
			if lit || (!known && infinity) {
				n++
			}
		}
		return n
	}

	for i := 0; i < 50; i++ {
		newpic := make(map[image.Point]bool)
		for p := range pic {
			for _, o := range offsets {
				q := p.Add(o)
				newpic[q] = program[index(q, i%2 == 1)] == '#'
			}
		}
		pic = newpic
		//printPic(pic)
	}

	count := 0
	for _, lit := range pic {
		if lit {
			count++
		}
	}
	fmt.Println(count)
}

func printPic(pic map[image.Point]bool) {
	var bounds image.Rectangle
	for p := range pic {
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

	out := make([][]byte, bounds.Max.Y-bounds.Min.Y+1)
	for j := range out {
		out[j] = bytes.Repeat([]byte{'.'}, bounds.Max.X-bounds.Min.X+1)
	}

	for p, lit := range pic {
		if lit {
			out[p.Y-bounds.Min.Y][p.X-bounds.Min.X] = '#'
		}
	}

	for _, line := range out {
		fmt.Println(string(line))
	}
}
