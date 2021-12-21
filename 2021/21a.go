package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	p1pos, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid start for player 1: %v", err)
	}
	p2pos, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid start for player 2: %v", err)
	}

	ddie := 1
	rolls := 0
	p1score, p2score := 0, 0
	p1turn := true
	for p1score < 1000 && p2score < 1000 {
		if p1turn {
			p1pos = (p1pos+3*(ddie+1)-1)%10 + 1
			p1score += p1pos
		} else {
			p2pos = (p2pos+3*(ddie+1)-1)%10 + 1
			p2score += p2pos
		}
		ddie = (ddie+2)%100 + 1
		rolls += 3
		p1turn = !p1turn
	}

	min := p1score
	if p2score < min {
		min = p2score
	}
	fmt.Println(rolls * min)
}
