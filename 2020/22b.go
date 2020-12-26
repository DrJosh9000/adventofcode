package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func atoi(x string) int {
	n, err := strconv.Atoi(x)
	if err != nil {
		die("Couldn't parse %q: %v", x, err)
	}
	return n
}

func score(x []byte) int {
	s := 0
	for i, c := range x {
		s += int(c) * (len(x) - i)
	}
	return s
}

func combat(a, b []byte) (awins bool, sc int) {
	type state struct{ a, b string }
	seen := make(map[state]struct{})
	for len(a) > 0 && len(b) > 0 {
		// Check infinite game rule.
		st := state{string(a), string(b)}
		if _, dejavu := seen[st]; dejavu {
			// Player 1 wins this *game*
			return true, score(a)
		}
		seen[st] = struct{}{}

		// If recurisng is possible, recurse...
		if a[0] < byte(len(a)) && b[0] < byte(len(b)) {
			a0, b0 := make([]byte, a[0]), make([]byte, b[0])
			copy(a0, a[1:])
			copy(b0, b[1:])
			awins, _ := combat(a0, b0)
			if awins {
				// Player 1 wins the round
				a = append(a[1:], a[0], b[0])
				b = b[1:]
			} else {
				// Player 2 wins the round
				b = append(b[1:], b[0], a[0])
				a = a[1:]
			}
			continue
		}

		// Highest card wins the round
		if a[0] > b[0] {
			// Player 1 wins the round
			a = append(a[1:], a[0], b[0])
			b = b[1:]
		} else {
			// Player 2 wins the round
			b = append(b[1:], b[0], a[0])
			a = a[1:]
		}
	}
	// Whoever has cards left wins
	if len(a) > 0 {
		return true, score(a)
	}
	return false, score(b)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	var a, b []byte
	deck := &a
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		switch m {
		case "Player 1:", "":
			// skip
		case "Player 2:":
			deck = &b
		default:
			*deck = append(*deck, byte(atoi(m)))
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	_, score := combat(a, b)
	fmt.Println(score)
}
