package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	type coord struct{ x, y int }
	tiles := make(map[coord]struct{})
	//     (-1,1)   (0,1)
	// (-1,0)   (0,0)   (1,0)
	//     (0,-1)   (1,-1)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		var c coord
		for m != "" {
			switch {
			case strings.HasPrefix(m, "e"):
				c.x++
				m = m[1:]
			case strings.HasPrefix(m, "w"):
				c.x--
				m = m[1:]
			case strings.HasPrefix(m, "ne"):
				c.y++
				m = m[2:]
			case strings.HasPrefix(m, "nw"):
				c.x--
				c.y++
				m = m[2:]
			case strings.HasPrefix(m, "se"):
				c.x++
				c.y--
				m = m[2:]
			case strings.HasPrefix(m, "sw"):
				c.y--
				m = m[2:]
			}
		}
		if _, black := tiles[c]; black {
			delete(tiles, c)
		} else {
			tiles[c] = struct{}{}
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(len(tiles))
}
