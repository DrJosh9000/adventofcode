package main

import (
	"fmt"
	"image"
	"log"

	"github.com/DrJosh9000/exp"
)

func main() {
	var tubes []string
	exp.MustForEachLineIn("inputs/19.txt", func(line string) {
		tubes = append(tubes, line)
	})
	
	startx := -1
	for i, c := range tubes[0] {
		if c == '|' {
			startx = i
			break
		}
	}
	if startx == -1 {
		log.Fatal("Couldn't find start position")
	}
	
	steps := 0
	p := image.Pt(startx, 0)
	d := image.Pt(0, 1)
	var letters []byte
loop:
	for {
		c := tubes[p.Y][p.X]
		switch c {
		case ' ':
			break loop
		case '-', '|':
			// This case includes cross-overs
		case '+':
			nd1, nd2 := image.Pt(-d.Y, d.X), image.Pt(d.Y, -d.X)
			if p1 := p.Add(nd1); tubes[p1.Y][p1.X] != ' ' {
				d = nd1
			} else {
				d = nd2
			}
		default:
			letters = append(letters, c)
		}
		p = p.Add(d)
		steps++
	}
		
	fmt.Printf("%s\n%d steps\n", string(letters), steps)
}