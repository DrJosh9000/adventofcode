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

	h, v := 0, 0
	for _, s := range steps {
		switch s.dir {
		case "forward":
			h += s.count
		case "down":
			v += s.count
		case "up":
			v -= s.count
		}
	}
	fmt.Println(h * v)
}
