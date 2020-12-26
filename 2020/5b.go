package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var ids []int
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
		ids = append(ids, int(id))
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	sort.Ints(ids)
	for i, id := range ids[1:] {
		if id-ids[i] == 2 {
			fmt.Println(id-1)
			return
		}
	}
}
