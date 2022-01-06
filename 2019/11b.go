package main

import (
	"image"
	"log"

	"github.com/DrJosh9000/adventofcode/2019/intcode"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	vm := intcode.ReadProgram("inputs/11.txt")
	in, out := make(chan int), make(chan int)
	go vm.Run(in, out)

	hull := map[image.Point]int{{}: 1}
	p, o := image.Pt(0, 0), image.Pt(0, 1)
commLoop:
	for {
		select {
		case c, ok := <-out:
			if !ok {
				// done
				break commLoop
			}
			hull[p] = c
			if <-out == 0 {
				o = image.Pt(-o.Y, o.X)
			} else {
				o = image.Pt(o.Y, -o.X)
			}
			p = p.Add(o)
		case in <- hull[p]:
			// nop
		}
	}

	pts := make(plotter.XYs, 0, len(hull))
	for p, x := range hull {
		if x == 0 {
			continue
		}
		pts = append(pts, plotter.XY{float64(p.X), float64(p.Y)})
	}
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatalf("Couldn't create scatter plot: %v", err)
	}
	scatter.Shape = draw.BoxGlyph{}
	pl := plot.New()
	pl.Add(scatter)
	if err := pl.Save(9*vg.Centimeter, 2*vg.Centimeter, "11b.png"); err != nil {
		log.Fatalf("Couldn't save plot: %v", err)
	}
}
