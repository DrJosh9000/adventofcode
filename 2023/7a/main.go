package main

import (
	"fmt"
	"strings"

	"github.com/DrJosh9000/exp"
	"github.com/DrJosh9000/exp/algo"
)

// Advent of Code 2023
// Day 7, part a

const cards = "23456789TJQKA"

const (
	highCard = iota
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

func handType(hand string) int {
	hist := algo.Freq([]rune(hand))
	switch len(hist) {
	case 1:
		return fiveOfKind

	case 2:
		for _, n := range hist {
			if n == 1 || n == 4 {
				return fourOfKind
			}
			return fullHouse
		}

	case 3:
		for _, n := range hist {
			if n == 3 {
				return threeOfKind
			}
			if n == 2 {
				return twoPair
			}
		}

	case 4:
		return onePair

	case 5:
		return highCard
	}
	panic("couldn't classify hand " + hand)
}

func main() {
	type handBid struct {
		hand string
		bid  int
	}

	var input []handBid

	for _, line := range exp.MustReadLines("2023/inputs/7.txt") {
		var hand handBid
		exp.Must(fmt.Sscanf(line, "%s %d", &hand.hand, &hand.bid))
		input = append(input, hand)
	}

	algo.SortSlice(input, func(a, b handBid) bool {
		at, bt := handType(a.hand), handType(b.hand)
		if at != bt {
			return at < bt
		}
		for k := 0; k < 5; k++ {
			ac, bc := a.hand[k], b.hand[k]
			if ac == bc {
				continue
			}
			return strings.IndexByte(cards, ac) < strings.IndexByte(cards, bc)
		}
		return false
	})

	sum := 0
	for i, h := range input {
		sum += (i + 1) * h.bid
	}

	fmt.Println(sum)
}
