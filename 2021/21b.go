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

	const goal = 21

	// at each turn the pawns can be in one of 100 states.
	// the players each have a score, and it can be either p1's turn or p2's.
	type state struct {
		p1pos, p2pos     int
		p1score, p2score int
		p2turn           bool
	}
	// start with the one initial universe
	states := map[state]int{
		{
			p1pos:   p1pos,
			p2pos:   p2pos,
			p1score: 0,
			p2score: 0,
			p2turn:  false,
		}: 1,
	}

	// splits[n] is number of universes created for a roll of n.
	splits := []int{0, 0, 0, 1, 3, 6, 7, 6, 3, 1}

	// let's play dirac dice
	p1wins, p2wins := 0, 0
	for len(states) > 0 {
		for s, u := range states {
			delete(states, s)
			// there are u universes in state s
			// each splits into multiple universes
			if !s.p2turn {
				// p1's turn
				for n, d := range splits {
					if n == 0 {
						continue
					}
					pos := (s.p1pos+n-1)%10 + 1
					score := s.p1score + pos
					if score >= goal {
						p1wins += d * u
						continue
					}
					states[state{
						p1pos:   pos,
						p2pos:   s.p2pos,
						p1score: score,
						p2score: s.p2score,
						p2turn:  true,
					}] += d * u
				}
			} else {
				// p2's turn
				for n, d := range splits {
					if n == 0 {
						continue
					}
					pos := (s.p2pos+n-1)%10 + 1
					score := s.p2score + pos
					if score >= goal {
						p2wins += d * u
						continue
					}
					states[state{
						p1pos:   s.p1pos,
						p2pos:   pos,
						p1score: s.p1score,
						p2score: score,
						p2turn:  false,
					}] += d * u
				}
			}
		}
	}

	fmt.Println(p1wins)
	fmt.Println(p2wins)
}
