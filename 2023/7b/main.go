package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"drjosh.dev/exp"
	"drjosh.dev/exp/algo"
)

// Advent of Code 2023
// Day 7, part b

const cards = "J23456789TQKA"

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
		if hist['J'] > 0 {
			return fiveOfKind
		}

		for _, n := range hist {
			if n == 1 || n == 4 {
				return fourOfKind
			}

			return fullHouse
		}

	case 3:
		switch hist['J'] {
		case 2, 3:
			return fourOfKind
		case 1:
			for c, n := range hist {
				if c == 'J' {
					continue
				}
				if n == 3 {
					return fourOfKind
				}
				if n == 2 {
					return fullHouse
				}
			}
		}

		for _, n := range hist {
			if n == 3 {
				return threeOfKind
			}
			if n == 2 {
				return twoPair
			}
		}

	case 4:
		if hist['J'] > 0 {
			return threeOfKind
		}
		return onePair

	case 5:
		if hist['J'] > 0 {
			return onePair
		}
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

	slices.SortFunc(input, func(a, b handBid) int {
		at, bt := handType(a.hand), handType(b.hand)
		if at != bt {
			return cmp.Compare(at, bt)
		}
		for k := 0; k < 5; k++ {
			ac, bc := a.hand[k], b.hand[k]
			if ac == bc {
				continue
			}
			return cmp.Compare(strings.IndexByte(cards, ac), strings.IndexByte(cards, bc))
		}
		return 0
	})

	sum := 0
	for i, h := range input {
		sum += (i + 1) * h.bid
	}

	fmt.Println(sum)
}
