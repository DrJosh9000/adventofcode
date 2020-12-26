package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lineRE = regexp.MustCompile(`^(\d+)-(\d+) (.): ([a-z]+)$`)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.2")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	valid := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := lineRE.FindStringSubmatch(sc.Text())
		if len(m) != 5 {
			die("Line did not match regexp: %q !~ %v", sc.Text(), lineRE)
		}
		l, err := strconv.Atoi(m[1])
		if err != nil {
			die("Couldn't parse %q: %v", m[1], err)	
		}
		h, err := strconv.Atoi(m[2])
		if err != nil {
			die("Couldn't parse %q: %v", m[2], err)
		}
		c := m[3][0]
		p := m[4]

		if (p[l-1] == c) != (p[h-1] == c) {
			valid++
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(valid)
}
