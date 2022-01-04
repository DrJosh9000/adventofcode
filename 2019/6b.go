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
	pred := make(map[string]string)
	var dfs func(string, string, int)
	dfs = func(v, p string, d int) {
		depth[v] = d
		pred[v] = p
		for _, w := range g[v] {
			dfs(w, v, d+1)
		}
	}
	dfs("COM", "", 0)
	x := "YOU"
	for x != "" {
		x = pred[x]
		y := "SAN"
		for y != "" {
			y = pred[y]
			if x == y {
				fmt.Println(depth["YOU"] + depth["SAN"] - 2*depth[x] - 2)
				return
			}
		}
	}
}
