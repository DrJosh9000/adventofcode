package main

import (
	"fmt"
	"io"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.11")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	var m [][]byte
	for {
		var line string
		_, err := fmt.Fscanf(f, "%s\n", &line)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		m = append(m, []byte(line))
	}
	w, h := len(m[0]), len(m)

	type offset struct{ x, y int }
	offs := []offset{
		{-1,-1}, {0,-1}, {1,-1},
		{-1, 0},         {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	nm := make([][]byte, h)
	for j := range nm {
		nm[j] = make([]byte, w)
	}
	
	for it:=0; ; it++ {
		changed := 0
		for j := range m {
			for i, c := range m[j] {
				if c != 'L' && c != '#' {
					continue
				}
				occ := 0
				for _, d := range offs {
					for mul:=1; ; mul++ {
						x, y := i+d.x*mul, j+d.y*mul
						if x<0 || y<0 || x>=w || y>=h {
							break
						}
						if m[y][x] == '#' {
							occ++
							break
						}
						if m[y][x] == 'L' {
							break
						}
					}
				}
				nm[j][i] = c		
				if c == '#' && occ >= 5 {
					nm[j][i] = 'L'
					changed++
				} 
				if c == 'L' && occ == 0 {
					nm[j][i] = '#'
					changed++
				}
			}
		}
		if changed == 0 {
			break
		}
		fmt.Printf("iteration %d changed %d\n", it, changed)
		m, nm = nm, m
	}
	count := 0
	for _, r := range m {
		for _, c := range r {
			if c == '#' {
				count++
			}
		}
	}  
	fmt.Println(count)
}
