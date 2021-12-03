package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("2.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	type step struct {
		dir   string
		count int
	}
	var steps []step
	for {
		var s step
		_, err := fmt.Fscan(f, &s.dir, &s.count)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't scan: %v", err)
		}
		steps = append(steps, s)
	}

	h, v, aim := 0, 0, 0
	for _, s := range steps {
		switch s.dir {
		case "forward":
			h += s.count
			v += aim * s.count
		case "down":
			aim += s.count
		case "up":
			aim -= s.count
		}
	}
	fmt.Println(h * v)
}
