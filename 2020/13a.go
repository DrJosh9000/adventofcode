package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.13")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	var arrival int
	var rawids string
	fmt.Fscanf(f, "%d\n%s\n", &arrival, &rawids)

	min := 1 << 31
	var minid int
	for _, t := range strings.Split(rawids, ",") {
		if t == "x" {
			continue
		}
		id, err := strconv.Atoi(t)
		if err != nil {
			die("Couldn't atoi %q: %v", t, err)
		}

		wait := (arrival/id+1)*id - arrival
		if wait < min {
			min, minid = wait, id
		}
	}
	fmt.Println(min * minid)
}
