package main

import (
	"fmt"
	"os"
	"strings"

	"drjosh.dev/exp"
)

func main() {
	var line, rev [16]byte
	for i := range line {
		line[i] = byte(i)
		rev[i] = byte(i)
	}

	steps := strings.Split(strings.TrimSpace(string(exp.Must(os.ReadFile("inputs/16.txt")))), ",")

	for _, step := range steps {
		switch step[0] {
		case 's':
			var n int
			exp.Must(fmt.Sscanf(step, "s%d", &n))
			var l [16]byte
			for i := range line {
				l[(i+n)%16] = line[i]
			}
			line = l
			for i, c := range line {
				rev[c] = byte(i)
			}
		case 'x':
			var a, b int
			exp.Must(fmt.Sscanf(step, "x%d/%d", &a, &b))
			line[a], line[b] = line[b], line[a]
			rev[line[a]], rev[line[b]] = byte(a), byte(b)
		case 'p':
			var a, b rune
			exp.Must(fmt.Sscanf(step, "p%c/%c", &a, &b))
			i, j := rev[a-'a'], rev[b-'a']
			line[i], line[j] = line[j], line[i]
			rev[line[i]], rev[line[j]] = i, j
		}
	}

	for _, c := range line {
		fmt.Printf("%c", c+'a')
	}
	fmt.Println()
}
