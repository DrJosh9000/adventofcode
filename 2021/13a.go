package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/13.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	dots := make(map[image.Point]struct{})
	type fold struct {
		axis rune
		ord  int
	}
	var folds []fold
	sc := bufio.NewScanner(f)
	foldstate := false
	for sc.Scan() {
		switch t := sc.Text(); {
		case t == "":
			foldstate = true
			continue
		case foldstate:
			var f fold
			if _, err := fmt.Sscanf(sc.Text(), "fold along %c=%d", &f.axis, &f.ord); err != nil {
				log.Fatalf("Malformed fold: %v", err)
			}
			folds = append(folds, f)
		default:
			var p image.Point
			if _, err := fmt.Sscanf(sc.Text(), "%d,%d", &p.X, &p.Y); err != nil {
				log.Fatalf("Malformed point: %v", err)
			}
			dots[p] = struct{}{}
		}

	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	switch f := folds[0]; f.axis {
	case 'x':
		for p := range dots {
			if p.X > f.ord {
				delete(dots, p)
				p.X = 2*f.ord - p.X
				dots[p] = struct{}{}
			}
		}
	case 'y':
		for p := range dots {
			if p.Y > f.ord {
				delete(dots, p)
				p.Y = 2*f.ord - p.Y
				dots[p] = struct{}{}
			}
		}
	}

	fmt.Println(len(dots))

}
