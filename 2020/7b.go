package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var (
	ruleRE = regexp.MustCompile(`^(\w+ \w+) bags contain (.+)\.$`)
	contRE = regexp.MustCompile(`^(\d+) (\w+ \w+) bags?$`)
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.7")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	type bagspec struct{
		n int
		c string
	}
	conts := make(map[string][]bagspec)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		rule := ruleRE.FindStringSubmatch(m)
		if len(rule) != 3 {
			die("Couldn't parse rule %q", rule)
		}
		if rule[2] == "no other bags" {
			continue
		}
		for _, cont := range strings.Split(rule[2], ", ") {
			con := contRE.FindStringSubmatch(cont)
			if len(con) != 3 {
				die("Couldn't parse contained %q", cont)
			}
			if con[1] == "0" {
				continue
			}
			n, err := strconv.Atoi(con[1])
			if err != nil {
				die("Couldn't parse %q: %v", con[1], err)
			}
			conts[rule[1]] = append(conts[rule[1]], bagspec{
				n: n,
				c: con[2],
			})
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}

	memo := make(map[string]int)
	var count func(string) int
	count = func(c string) int {
		if n, ok := memo[c]; ok {
			return n
		}
		sum := 1 // this bag
		for _, bs := range conts[c] {
			sum += bs.n * count(bs.c)	
		}
		memo[c] = sum
		return sum
	}
	
	fmt.Println(count("shiny gold")-1)
}
