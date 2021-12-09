package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/9.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var heightmap []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		heightmap = append(heightmap, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	risk := 0
	for y, row := range heightmap {
		for x := range row {
			cell := row[x]
			if x > 0 && cell >= row[x-1] {
				continue
			}
			if x < len(row)-1 && cell >= row[x+1] {
				continue
			}
			if y > 0 && cell >= heightmap[y-1][x] {
				continue
			}
			if y < len(heightmap)-1 && cell >= heightmap[y+1][x] {
				continue
			}
			risk += int(cell-'0') + 1
		}
	}

	fmt.Println(risk)
}
