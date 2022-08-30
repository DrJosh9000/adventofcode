package main

import (
	"fmt"
	"strings"
	"os"

	"github.com/DrJosh9000/exp"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	path := strings.Split(strings.TrimSpace(string(exp.Must(os.ReadFile("inputs/11.txt")))), ",")
	
	n, ne, nw := 0, 0, 0
	for _, step := range path {
		switch step {
		case "n": n++
		case "s": n--
		case "ne": ne++
		case "sw": ne--
		case "nw": nw++
		case "se": nw--
		}
	}
	
	// normalise ne,nw and se,sw
	switch {
	case ne > 0 && nw > 0: 	
		// nw,ne = ne,nw = n
		m := min(ne, nw)
		n += m
		ne -= m
		nw -= m
	case ne < 0 && nw < 0:	
		// sw,se = se,sw = s
		m := min(abs(ne), abs(nw))
		n -= m
		ne += m
		nw += m
	}
	
	// normalise n,s{e,w} and s,n{e,w}
	// (nb: ne,se faster than n,se,se)
	switch {
	case n < 0 && ne > 0:
		// s,ne = ne,s = se
		m := min(abs(n), ne)
		n += m
		ne -= m
		nw -= m
	case n < 0 && nw > 0:
		// s,nw = nw,s = sw	
		m := min(abs(n), nw)
		n += m
		nw -= m
		ne -= m
	case n > 0 && ne < 0:
		// n,sw = sw,n = nw
		m := min(n, abs(ne))
		n -= m
		ne += m
		nw += m
	case n > 0 && nw < 0:
		// n,se = se,n = ne
		m := min(n, abs(nw))
		n -= m
		nw += m
		ne += m
	}
		
	fmt.Println(abs(n) + abs(ne) + abs(nw))
}