package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/25.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}
	defer f.Close()

	var input []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan input: %v", err)
	}

	seafloor := make([][]rune, len(input))
	for j := range input {
		seafloor[j] = make([]rune, len(input[j]))
		for i, c := range input[j] {
			seafloor[j][i] = c
		}
	}
	h, w := len(seafloor), len(seafloor[0])
	steps := 0
	for {
		moved := false

		// find > sea cucumbers to move
		move := make(map[image.Point]struct{})
		for j, row := range seafloor {
			for i, c := range row {
				if c == '>' && row[(i+1)%w] == '.' {
					moved = true
					move[image.Pt(i, j)] = struct{}{}
				}
			}
		}
		// move them
		for p := range move {
			seafloor[p.Y][p.X] = '.'
			seafloor[p.Y][(p.X+1)%w] = '>'
			delete(move, p)
		}

		// now do it all again for the v cucumbers
		for j := range seafloor {
			for i, c := range seafloor[j] {
				if c == 'v' && seafloor[(j+1)%h][i] == '.' {
					moved = true
					move[image.Pt(i, j)] = struct{}{}
				}
			}
		}
		for p := range move {
			seafloor[p.Y][p.X] = '.'
			seafloor[(p.Y+1)%h][p.X] = 'v'
		}
		steps++
		/*
			fmt.Println("After step", steps)
			for j := range seafloor {
				for _, c := range seafloor[j] {
					fmt.Printf("%c", c)
				}
				fmt.Println()
			}
			if steps == 58 {
				break
			}
		*/
		// did any move?
		if !moved {
			break
		}
	}
	fmt.Println(steps)

}
