package main

import (
	"bufio"
	"fmt"
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
	}
	var rules []rule
	invalidSum := 0
	sc := bufio.NewScanner(f)
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
			r := rule{
				name: rs[1],
				al:   atoi(rs[2]),
				ah:   atoi(rs[3]),
				bl:   atoi(rs[4]),
				bh:   atoi(rs[5]),
			}
			rules = append(rules, r)

		case stMine:
			if m == "nearby tickets:" {
				state = stNearby
				continue
			}
			// must be my ticket
			// skip for now
		case stNearby:
			// must be a nearby ticket
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
					invalidSum += n
				}
			}
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(invalidSum)
}
