package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/8.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var lines [][]string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, strings.Fields(sc.Text()))
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	count := 0
	for _, line := range lines {
		for _, token := range line[11:] {
			switch len(token) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	fmt.Println(count)
}
