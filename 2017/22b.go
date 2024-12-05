package main

import (
	"fmt"
	"image"

	"drjosh.dev/exp"
)

func main() {
	grid := make(map[image.Point]byte)
	y := 0
	pos := image.Pt(0, 0)
	exp.MustForEachLineIn("inputs/22.txt", func(line string) {
		pos.X = len(line) / 2
		for x, c := range line {
			grid[image.Pt(x, y)] = byte(c)
		}
		y++
	})
	pos.Y = y / 2
	dir := image.Pt(0, -1)
	count := 0
	for burst := 0; burst < 10_000_000; burst++ {
		switch grid[pos] {
		case 0, '.':
			dir = image.Pt(dir.Y, -dir.X)
			grid[pos] = 'W'
		case 'W':
			// does not turn
			grid[pos] = '#'
			count++
		case '#':
			dir = image.Pt(-dir.Y, dir.X)
			grid[pos] = 'F'
		case 'F':
			dir = image.Pt(-dir.X, -dir.Y)
			grid[pos] = '.'
		}
		pos = pos.Add(dir)
	}
	fmt.Println(count)
}
