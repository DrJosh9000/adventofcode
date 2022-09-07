package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"os"

	"github.com/DrJosh9000/exp"
)

type grid2[T any] [2][2]T

// rot rotates the grid clockwise
func (g grid2[T]) rot() grid2[T] {
	return grid2[T]{
		{g[1][0], g[0][0]},
		{g[1][1], g[0][1]},
	}
}

// flip flips the grid horizontally
func (g grid2[T]) flip() grid2[T] {
	return grid2[T]{
		{g[0][1], g[0][0]},
		{g[1][1], g[1][0]},
	}
}

type grid3[T any] [3][3]T

// rot rotates the grid clockwise
func (g grid3[T]) rot() grid3[T] {
	return grid3[T]{
		{g[2][0], g[1][0], g[0][0]},
		{g[2][1], g[1][1], g[0][1]},
		{g[2][2], g[1][2], g[0][2]},
	}
}

// flip flips the grid horizontally
func (g grid3[T]) flip() grid3[T] {
	return grid3[T]{
		{g[0][2], g[0][1], g[0][0]},
		{g[1][2], g[1][1], g[1][0]},
		{g[2][2], g[2][1], g[2][0]},
	}
}

func main() {	
	rules := make(map[any][]string)
	exp.MustForEachLineIn("inputs/21.txt", func(line string) {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			log.Fatalf("len(parts) = %d [want 2]", len(parts))
		}
		match := strings.Split(parts[0], "/")
		output := strings.Split(parts[1], "/")
		switch len(match) {
		case 2:
			var g grid2[byte]
			for i := range g {
				copy(g[i][:], match[i])
			}
			for h := 0; h < 2; h++ {
				for r := 0; r < 4; r++ {
					rules[g] = output
					g = g.rot()
				}
				g = g.flip()
			}
		case 3:
			var g grid3[byte]
			for i := range g {
				copy(g[i][:], match[i])
			}
			for h := 0; h < 2; h++ {
				for r := 0; r < 4; r++ {
					rules[g] = output
					g = g.rot()
				}
				g = g.flip()
			}
		default:
			log.Fatalf("len(match) = %d [want 2 or 3]", len(match))
		}
	})
	
	grid := [][]byte{
		[]byte(".#."),
		[]byte("..#"),
		[]byte("###"),
	}
	
	for i := 0; i < exp.Must(strconv.Atoi(os.Args[1])); i++ {
		var cs, ncs int
		if len(grid) % 2 == 0 {
			cs, ncs = 2, 3
		} else {
			cs, ncs = 3, 4
		}
		ns := len(grid) / cs * ncs
		ng := make([][]byte, ns)
		for j := range ng {
			ng[j] = make([]byte, ns)
		}
		
		for r := 0; r < len(grid)/cs; r++ {
			for c := 0; c < len(grid)/cs; c++ {
				switch cs {
				case 2:
					var g grid2[byte]
					for j := range g {
						copy(g[j][:], grid[r*cs+j][c*cs:])
					}
					for j, o := range rules[g] {
						copy(ng[r*ncs+j][c*ncs:], o)
					}
				case 3:
					var g grid3[byte]
					for j := range g {
						copy(g[j][:], grid[r*cs+j][c*cs:(c+1)*cs])
					}
					for j, o := range rules[g] {
						copy(ng[r*ncs+j][c*ncs:], o)
					}
				}
			}
		}
		grid = ng
	}
	
	count := 0
	for _, r := range grid {
		for _, c := range r {
			if c == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}
