package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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

	cby := make(map[string][]string)
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
			cby[con[2]] = append(cby[con[2]], rule[1])
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	gold := make(map[string]struct{})
	queue := []string{"shiny gold"}
	bag := ""
	for len(queue) > 0 {
		bag, queue = queue[0], queue[1:]
		for _, b := range cby[bag] {
			if _, seen := gold[b]; seen {
				continue
			}
			gold[b] = struct{}{}
			queue = append(queue, b)
		}
	}
	fmt.Println(len(gold))
}
