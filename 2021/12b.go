package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("inputs/12.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	graph := make(map[string][]string)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		vs := strings.Split(sc.Text(), "-")
		if len(vs) < 2 {
			log.Fatalf("Malformed line: %q", sc.Text())
		}
		u, v := vs[0], vs[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	paths := 0
	var dfs func([]string, string)
	dfs = func(path []string, double string) {
		u := path[len(path)-1]
		if u == "end" {
			paths++
			return
		}
	destLoop:
		for _, v := range graph[u] {
			if unicode.IsUpper(rune(v[0])) {
				dfs(append(path, v), double)
				continue
			}
			if v == "start" {
				continue
			}
			for _, p := range path {
				if v == p {
					if double == "" {
						dfs(append(path, v), v)
					}
					continue destLoop
				}
			}
			dfs(append(path, v), double)
		}
	}
	dfs([]string{"start"}, "")

	fmt.Println(paths)
}
