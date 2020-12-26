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

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	var a, b []int
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
			*deck = append(*deck, atoi(m))
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			// Player 1 wins
			a = append(a[1:], a[0], b[0])
			b = b[1:]
		} else {
			// Player 2 wins
			b = append(b[1:], b[0], a[0])
			a = a[1:]
		}
	}
	// find the winner
	winner := a
	if len(a) == 0 {
		winner = b
	}
	sum := 0
	for i, c := range winner {
		sum += c * (len(winner) - i)
	}
	fmt.Println(sum)
}
