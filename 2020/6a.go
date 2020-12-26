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
	ques := make(map[rune]struct{})
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		if m == "" {
			total += len(ques)
			ques = make(map[rune]struct{})
			continue
		}
		for _, r := range m {
			ques[r] = struct{}{}
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	total += len(ques)
	fmt.Println(total)
}
