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

	m := make([][]byte, 1)
	for {
		var line string
		_, err := fmt.Fscanf(f, "%s\n", &line)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		m = append(m, []byte("." + line + "."))
	}
	w := len(m[1])
	m[0] = make([]byte, w) 
	m = append(m, make([]byte, w))
	h := len(m)

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
		for i:=1; i<w-1; i++ {
			for j:=1; j<h-1; j++ {
				if m[j][i] != 'L' && m[j][i] != '#' {
					continue
				}
				occ := 0
				for _, d := range offs {
					if m[j+d.y][i+d.x] == '#' {
						occ++
					}
				}
				nm[j][i] = m[j][i]		
				if m[j][i] == '#' && occ >= 4 {
					nm[j][i] = 'L'
					changed++
				} 
				if m[j][i] == 'L' && occ == 0 {
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
