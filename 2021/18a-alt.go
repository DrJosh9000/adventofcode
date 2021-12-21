package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var bigNumberRE = regexp.MustCompile(`\d{2,}`)

type any = interface{}

func main() {
	f, err := os.Open("inputs/18.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()

	var numbers []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		numbers = append(numbers, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Couldn't sc.scan: %v", err)
	}

	x := numbers[0]
	for _, y := range numbers[1:] {
		x = add(x, y)
	}
	var expr []any
	if err := json.Unmarshal([]byte(x), &expr); err != nil {
		log.Fatalf("Couldn't unmarshal: %v", err)
	}
	fmt.Println(magnitude(expr))
}

// magnitude computes the magnitude of a pair or int.
func magnitude(x any) int {
	switch x := x.(type) {
	case float64:
		return int(x)
	case []any:
		return 3*magnitude(x[0]) + 2*magnitude(x[1])
	}
	panic("invalid type in expression")
}

func add(x, y string) string {
	return reduce(fmt.Sprintf("[%s,%s]", x, y))
}

// reduce applies reductions (explode or split).
func reduce(x string) string {
	for {
		if y, did := explode(x); did {
			x = y
			continue
		}
		if y, did := split(x); did {
			x = y
			continue
		}
		break
	}
	return x
}

// explode explodes the first pair nested at least 4 deep.
func explode(x string) (string, bool) {
	depth := 0
	for i, r := range x {
		switch r {
		case '[':
			depth++
			if depth < 5 {
				continue
			}
			var l, r int
			if _, err := fmt.Sscanf(x[i+1:], "%d,%d", &l, &r); err != nil {
				log.Fatalf("Couldn't read pair: %v (input %q)", err, x[i+1:])
			}
			// look backwards
			first, last := -1, -1
			for j := i; j >= 0; j-- {
				if last == -1 && unicode.IsDigit(rune(x[j])) {
					last = j + 1
					continue
				}
				if last != -1 && !unicode.IsDigit(rune(x[j])) {
					first = j + 1
					break
				}
			}
			y := x[:i]
			if first != -1 && last != -1 {
				n, err := strconv.Atoi(x[first:last])
				if err != nil {
					log.Fatalf("Couldn't parse number: %v", err)
				}
				y = x[:first] + strconv.Itoa(n+l) + x[last:i]
			}
			y += "0"
			end := -1
			for j := i + 1; j < len(x); j++ {
				if x[j] == ']' {
					end = j
					break
				}
			}
			first, last = -1, -1
			for j := end + 1; j < len(x); j++ {
				if first == -1 && unicode.IsDigit(rune(x[j])) {
					first = j
					continue
				}
				if first != -1 && !unicode.IsDigit(rune(x[j])) {
					last = j
					break
				}
			}
			if first < 0 && last < 0 {
				y += x[end+1:]
				return y, true
			}
			n, err := strconv.Atoi(x[first:last])
			if err != nil {
				log.Fatalf("Couldn't parse number: %v", err)
			}
			y += x[end+1:first] + strconv.Itoa(n+r) + x[last:]
			return y, true

		case ']':
			depth--
		}
	}
	return x, false
}

// split splits the leftmost number larger than 9
func split(x string) (string, bool) {
	match := bigNumberRE.FindStringIndex(x)
	if match == nil {
		return x, false
	}
	n, err := strconv.Atoi(x[match[0]:match[1]])
	if err != nil {
		log.Fatalf("Couldn't parse number: %v", err)
	}
	l := n / 2
	r := n - l
	return x[:match[0]] + fmt.Sprintf("[%d,%d]", l, r) + x[match[1]:], true
}
