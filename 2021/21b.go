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
		p1turn           bool
	}
	// start with the one initial universe
	states := map[state]int{
		{
			p1pos:   p1pos,
			p2pos:   p2pos,
			p1score: 0,
			p2score: 0,
			p1turn:  true,
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
			for n, d := range splits {
				if n == 0 {
					continue
				}
				st := s
				if s.p1turn {
					st.p1pos = (s.p1pos+n-1)%10 + 1
					st.p1score += st.p1pos
					if st.p1score >= goal {
						p1wins += d * u
						continue
					}
				} else { // p2's turn
					st.p2pos = (s.p2pos+n-1)%10 + 1
					st.p2score += st.p2pos
					if st.p2score >= goal {
						p2wins += d * u
						continue
					}
				}
				st.p1turn = !st.p1turn
				states[st] += d * u
			}
		}
	}

	fmt.Println(p1wins)
	fmt.Println(p2wins)
}
