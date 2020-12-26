package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

func atoi(x string) int {
	n, err := strconv.Atoi(x)
	if err != nil {
		die("Couldn't parse %q: %v", x, err)
	}
	return n
}

type expr interface {
	eval() int
	String() string
}

type num int

func (n num) eval() int      { return int(n) }
func (n num) String() string { return strconv.Itoa(int(n)) }

type binop struct {
	l, r expr
	o    string
}

func (o *binop) eval() int {
	if o.o == "*" {
		return o.l.eval() * o.r.eval()
	}
	return o.l.eval() + o.r.eval()
}

func (o *binop) String() string {
	return fmt.Sprintf("(%s %s %s)", o.l, o.o, o.r)
}

func tokenise(e string) []string {
	e = strings.ReplaceAll(e, "(", "( ")
	e = strings.ReplaceAll(e, ")", " )")
	return strings.Split(e, " ")
}

func parse(tokens []string) expr {
	var x expr
	lev, b := 0, 0
	for i := range tokens {
		tok := tokens[i]
		switch tok {
		case "(":
			if lev == 0 {
				b = i + 1
			}
			lev++
		case ")":
			lev--
			if lev > 0 {
				continue
			}
			y := parse(tokens[b:i])
			if x == nil {
				x = y
			} else {
				x.(*binop).r = y
			}
		case "+", "*":
			if lev > 0 {
				continue
			}
			x = &binop{
				l: x,
				o: tok,
			}
		default:
			if lev > 0 {
				continue
			}
			n := num(atoi(tok))
			if x == nil {
				x = n
			} else {
				x.(*binop).r = n
			}
		}
	}
	return x
}

func main() {
	f, err := os.Open("input.18")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	sum := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		sum += parse(tokenise(m)).eval()
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(sum)
}
