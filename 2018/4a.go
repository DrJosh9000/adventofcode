package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"drjosh.dev/exp"
)

func main() {
	type record struct {
		t time.Time
		e string
	}
	var recs []record
	exp.MustForEachLineIn("inputs/4.txt", func(line string) {
		parts := strings.Split(line, "] ")
		if len(parts) != 2 {
			log.Fatalf("Malformed line %q", line)
		}
		t, err := time.Parse("[2006-01-02 15:04", parts[0])
		if err != nil {
			log.Fatalf("Parsing timestamp: %v", err)
		}
		recs = append(recs, record{
			t: t,
			e: parts[1],
		})
	})

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].t.Before(recs[j].t)
	})

	type gm struct {
		guard, minute int
	}
	sleep := make(map[int]int)
	chart := make(map[gm]int)
	var guard, t int
	for _, r := range recs {
		switch {
		case r.e == "wakes up":
			sleep[guard] += r.t.Minute() - t
			for i := t; i < r.t.Minute(); i++ {
				chart[gm{guard: guard, minute: i}]++
			}

		case r.e == "falls asleep":
			t = r.t.Minute()

		default: // begins shift
			n, err := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(r.e, " begins shift"), "Guard #"))
			if err != nil {
				log.Fatalf("Malformed event: %v", err)
			}
			guard = n
		}
	}
	max := math.MinInt
	for g, s := range sleep {
		if s > max {
			max = s
			guard = g
		}
	}

	max = math.MinInt
	var best gm
	for m, x := range chart {
		if m.guard != guard {
			continue
		}
		if x > max {
			max = x
			best = m
		}
	}
	fmt.Println(best.guard, best.minute, best.guard*best.minute)
}
