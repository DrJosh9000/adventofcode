package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.5")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	highest := int64(-1)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		m = strings.Map(func (r rune) rune {
			if r == 'F' || r == 'L' {
				return '0'
			}
			if r == 'B' || r == 'R' {
				return '1'
			}
			return r
		}, m)
		id, err := strconv.ParseInt(m, 2, 32)
		if err != nil {
			die("Couldn't parse %q: %v", m, err)
		}
		if id > highest {
			highest = id
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(highest)
}
