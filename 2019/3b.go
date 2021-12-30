package main

import (
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("inputs/3.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}

	dir := map[byte]image.Point{
		'U': image.Pt(0, 1),
		'D': image.Pt(0, -1),
		'L': image.Pt(-1, 0),
		'R': image.Pt(1, 0),
	}
	dm := []map[image.Point]int{
		make(map[image.Point]int),
		make(map[image.Point]int),
	}
	for j, l := range strings.Split(string(f), "\n") {
		var p image.Point
		d := 0
		for _, token := range strings.Split(l, ",") {
			n, err := strconv.Atoi(token[1:])
			if err != nil {
				log.Fatalf("Couldn't atoi: %v", err)
			}
			δ := dir[token[0]]
			for i := 0; i < n; i++ {
				p = p.Add(δ)
				d++
				if _, exists := dm[j][p]; !exists {
					dm[j][p] = d
				}
			}
		}
	}

	md := math.MaxInt
	for p, d0 := range dm[0] {
		d1, found := dm[1][p]
		if !found {
			continue
		}
		if d := d0 + d1; d < md {
			md = d
		}

	}
	fmt.Println(md)
}
