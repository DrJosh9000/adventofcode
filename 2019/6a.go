package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/6.txt")
	if err != nil {
		log.Fatalf("Couldn't read input: %v", err)
	}
	defer f.Close()

	g := make(map[string][]string)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		t := strings.Split(sc.Text(), ")")
		g[t[0]] = append(g[t[0]], t[1])
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't scan: %v", err)
	}

	depth := make(map[string]int)
	var dfs func(string, int)
	dfs = func(v string, d int) {
		depth[v] = d
		for _, w := range g[v] {
			dfs(w, d+1)
		}
	}
	dfs("COM", 0)
	sum := 0
	for _, d := range depth {
		sum += d
	}
	fmt.Println(sum)
}
