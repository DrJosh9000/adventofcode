package main

import (
	"fmt"
	"image"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

func main() {
	grid := make(algo.Set[image.Point])
	y := 0
	pos := image.Pt(0, 0)
	exp.MustForEachLineIn("inputs/22.txt", func(line string) {
		pos.X = len(line) / 2
		for x, c := range line {
			if c == '#' {
				grid.Insert(image.Pt(x, y))
			}
		}
		y++
	})
	pos.Y = y / 2
	dir := image.Pt(0, -1)
	count := 0
	for burst := 0; burst < 10_000; burst++ {
		if grid.Contains(pos) {
			dir = image.Pt(-dir.Y, dir.X)
			delete(grid, pos)
		} else {
			dir = image.Pt(dir.Y, -dir.X)
			grid.Insert(pos)
			count++
		}
		pos = pos.Add(dir)
	}
	fmt.Println(count)
}
