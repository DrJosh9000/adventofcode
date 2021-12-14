package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/14.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()
	template := sc.Text()
	rules := make(map[string]string)
	for sc.Scan() {
		t := sc.Text()
		if t == "" {
			continue
		}
		ts := strings.Split(t, " -> ")
		rules[ts[0]] = ts[1]
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	for i := 0; i < 10; i++ {
		var out strings.Builder
		out.WriteByte(template[0])
		for j := range template[1:] {
			in := template[j : j+2]
			ins := rules[in]
			out.WriteString(ins)
			out.WriteByte(in[1])
		}
		template = out.String()
		if i < 4 {
			fmt.Println("After step %d: %s", i, template)
		}
	}

	sum := make(map[rune]int)
	for _, r := range template {
		sum[r]++
	}
	min, max := math.MaxInt, math.MinInt
	for _, n := range sum {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println(max - min)
}
