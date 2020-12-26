package main

import (
	"bufio"
	"fmt"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.3")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	trees, x, y := 0, 0, 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if y == 0 {
			// First line is not needed
			x, y = 3, 1
			continue
		}
		m := sc.Text()
		if m[x%len(m)] == '#' {
			trees++
		}
		x += 3
		y++
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(trees)
}
