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
	f, err := os.Open("input.6")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	total := 0
	ques := make(map[rune]int)
	peop := 0
	tally := func() {
		for r, c := range ques {
			if c == peop { 
				total++
			}
			ques[r] = 0
		}
		peop = 0
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		if m == "" {
			tally()			
			continue
		}
		peop++
		for _, r := range m {
			ques[r]++
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	tally()
	fmt.Println(total)
}
