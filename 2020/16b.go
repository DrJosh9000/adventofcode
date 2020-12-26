package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ruleRE = regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
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
	f, err := os.Open("input.16")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	const (
		stRules = iota
		stMine
		stNearby
	)
	state := stRules
	type rule struct {
		name           string
		al, ah, bl, bh int
		fields         uint
		field          int
	}
	var rules []*rule
	var tickets [][]int
	var mine []int
	sc := bufio.NewScanner(f)
scanLoop:
	for sc.Scan() {
		m := sc.Text()
		if m == "" {
			continue
		}
		switch state {
		case stRules:
			if m == "your ticket:" {
				state = stMine
				continue
			}
			// must be a rule
			rs := ruleRE.FindStringSubmatch(m)
			if len(rs) != 6 {
				die("Couldn't match ruleRE: %q !~ %v", m, ruleRE)
			}
			r := &rule{
				name:  rs[1],
				al:    atoi(rs[2]),
				ah:    atoi(rs[3]),
				bl:    atoi(rs[4]),
				bh:    atoi(rs[5]),
				field: -1,
			}
			rules = append(rules, r)

		case stMine:
			if m == "nearby tickets:" {
				state = stNearby
				continue
			}
			// must be my ticket
			for _, ns := range strings.Split(m, ",") {
				mine = append(mine, atoi(ns))
			}
		case stNearby:
			// must be a nearby ticket; requires validating from part A
			var tick []int
			for _, ns := range strings.Split(m, ",") {
				n := atoi(ns)
				valid := false
				for _, r := range rules {
					if r.al <= n && n <= r.ah {
						valid = true
						break
					}
					if r.bl <= n && n <= r.bh {
						valid = true
						break
					}
				}
				if !valid {
					continue scanLoop
				}
				tick = append(tick, n)
			}
			tickets = append(tickets, tick)
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	// There are R rules, and thus R fields in each ticket.
	// Assume any field could be any rule, and proceed by elimination.
	for _, r := range rules {
		r.fields = (1 << len(rules)) - 1
	}
	for _, tick := range tickets {
		for i, v := range tick {
			for _, r := range rules {
				if r.fields&(1<<i) == 0 {
					// already eliminated
					continue
				}
				if (r.al <= v && v <= r.ah) || (r.bl <= v && v <= r.bh) {
					// could still be this rule
					continue
				}
				r.fields &^= 1 << i
			}
		}
	}
	for {
		process := false
		// For each rule, see if it must be a particular field
		for _, r := range rules {
			if r.field > -1 {
				continue // already processed
			}
			if bits.OnesCount(r.fields) > 1 {
				continue // not specific enough yet
			}
			process = true
			r.field = bits.TrailingZeros(r.fields)
			// clear this bit in all other rules
			for _, r2 := range rules {
				if r == r2 {
					continue
				}
				r2.fields &^= r.fields
			}
		}
		if !process {
			break
		}
	}
	// Now... compute the product of the "departure" fields for my ticket.
	prod := 1
	for _, r := range rules {
		if strings.HasPrefix(r.name, "departure") {
			prod *= mine[r.field]
		}
	}
	fmt.Println(prod)
}
