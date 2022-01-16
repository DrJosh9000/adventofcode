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
	f, err := os.Open("inputs/22.txt")
	if err != nil {
		log.Fatalf("Couldn't open input: %v", err)
	}
	defer f.Close()

	const N = 10007

	// Combine the shuffle instructions into a single linear function.
	// card x ends up in position y
	// y = (mx + b) % N
	// The deck starts in order: m = 1, b = 0.
	m, b := 1, 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := sc.Text()
		switch {
		case t == "deal into new stack":
			// The deck is reversed; the current function is multiplied by -1
			// and shifted so that 0 becomes N-1.
			m = -m % N
			b = (-b - 1) % N

		case strings.HasPrefix(t, "deal with increment "):
			// The current function is multiplied by n.
			n, err := strconv.Atoi(strings.TrimPrefix(t, "deal with increment "))
			if err != nil {
				log.Fatalf("Invalid increment %q: %v", t, err)
			}
			m = (n * m) % N
			b = (n * b) % N

		case strings.HasPrefix(t, "cut "):
			// The current function is shifted by n.
			n, err := strconv.Atoi(strings.TrimPrefix(t, "cut "))
			if err != nil {
				log.Fatalf("Invalid cut %q: %v", t, err)
			}
			b = (b - n) % N
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan input: %v", err)
	}

	fmt.Println((2019*m + b) % N)
}
