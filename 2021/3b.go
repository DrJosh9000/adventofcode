package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/3.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var words []string
	for {
		var s string
		_, err := fmt.Fscan(f, &s)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Couldn't scan: %v", err)
		}
		words = append(words, s)
	}

	ox, co := filter(words, true), filter(words, false)
	o, err := strconv.ParseInt(ox, 2, 64)
	c, err := strconv.ParseInt(co, 2, 64)
	fmt.Println(o * c)
}

func filter(words []string, most bool) string {
	for i := 0; i < 12; i++ {
		one, zero := 0, 0
		for _, s := range words {
			if s[i] == '1' {
				one++
			} else {
				zero++
			}
		}
		w := make([]string, 0, len(words))
		var c byte
		if most {
			c = '1'
			if one < zero {
				c = '0'
			}
		} else {
			c = '0'
			if zero > one {
				c = '1'
			}
		}
		for _, s := range words {
			if s[i] == c {
				w = append(w, s)
			}
		}
		if len(w) == 1 {
			return w[0]
		}
		words = w
	}
	log.Fatal("Filter failure")
	return ""
}
