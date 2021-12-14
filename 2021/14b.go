package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/14.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()
	template := sc.Text()
	rules := make(map[string]string)
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			continue
		}
		ts := strings.Split(t, " -> ")
		rules[ts[0]] = ts[1]
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	elements := make(map[byte]int)
	for i := range template {
		elements[template[i]]++
	}
	pairs := make(map[string]int)
	for j := range template[1:] {
		pairs[template[j:j+2]]++
	}
	for i := 0; i < 40; i++ {
		newpairs := make(map[string]int)
		for p, n := range pairs {
			nu := rules[p]
			a, b := p[0:1]+nu, nu+p[1:2]
			newpairs[a] += n
			newpairs[b] += n
			elements[nu[0]] += n
		}
		pairs = newpairs
	}

	min, max := math.MaxInt, math.MinInt
	for _, n := range elements {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println(max - min)
}
