package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/4.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var draws []int
	sc := bufio.NewScanner(f)
	sc.Scan()
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}
	for _, x := range strings.Split(sc.Text(), ",") {
		n, err := strconv.Atoi(x)
		if err != nil {
			log.Fatalf("Couldn't convert string to int: %v", err)
		}
		draws = append(draws, n)
	}
	//log.Printf("draws: %v", draws)
	sc.Scan()
	var boards []board
	var b board
	var row int
	for sc.Scan() {
		if sc.Text() == "" {
			// new board
			boards = append(boards, b)
			b = board{}
			row = 0
			continue
		}
		if _, err := fmt.Sscanf(sc.Text(), "%d %d %d %d %d", &b[row][0], &b[row][1], &b[row][2], &b[row][3], &b[row][4]); err != nil {
			log.Fatalf("Couldn't sscanf: %v", err)

		}
		row++
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}
	boards = append(boards, b)

	//log.Printf("boards: %v", boards)

	// let's play bingo
	for _, n := range draws {
		for i := range boards {
			boards[i].mark(n)
			if boards[i].won() {
				fmt.Println(boards[i].sum() * n)
				return
			}
		}
	}
}

type board [5][5]int

func (b *board) mark(n int) {
	for i, row := range *b {
		for j, cell := range row {
			if cell == n {
				b[i][j] = -1
			}
		}
	}
}

func (b *board) won() bool {
	// check each row
	for _, row := range b {
		win := true
		for _, cell := range row {
			if cell != -1 {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}

	// check each column
	for j := range b[0] {
		win := true
		for i := range b {
			if cell := b[i][j]; cell != -1 {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	return false
}

func (b *board) sum() int {
	s := 0
	for _, row := range *b {
		for _, cell := range row {
			if cell == -1 {
				continue
			}
			s += cell
		}
	}
	return s
}
