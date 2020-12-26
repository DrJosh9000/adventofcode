package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

type rule struct {
	l, r []int // r nil if no or
	raw  string
}

type rules []rule

// This place is not a place of honor... no highly esteemed deed is commemorated here... nothing valued is here.
// What is here was dangerous and repulsive to us.
func (rs rules) regexpr(i int) (re string) {
	r := &rs[i]
	if r.raw != "" {
		return r.raw
	}
	defer func() { r.raw = re }()
	sb := new(strings.Builder)
	switch i {
	case 8:
		return fmt.Sprintf("(%s)+", rs.regexpr(42))
	case 11:
		// err, godawful hax
		fmt.Fprintf(sb, "((%s%s)", rs.regexpr(42), rs.regexpr(31))
		for i := 2; i < 20; i++ {
			fmt.Fprintf(sb, "|((%s){%d}(%s){%d})", rs.regexpr(42), i, rs.regexpr(31), i)
		}
		sb.WriteString(")")
		return sb.String()
	}
	if len(r.r) > 0 {
		sb.WriteString("((")
		for _, n := range r.l {
			sb.WriteString(rs.regexpr(n))
		}
		sb.WriteString(")|(")
		for _, n := range r.r {
			sb.WriteString(rs.regexpr(n))
		}
		sb.WriteString("))")
		return sb.String()
	}
	sb.WriteString("(")
	for _, n := range r.l {
		sb.WriteString(rs.regexpr(n))
	}
	sb.WriteString(")")
	return sb.String()
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	rules := make(rules, 150)
	var zeroRE *regexp.Regexp
	var messages bool
	valid := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		if m == "" {
			zeroRE = regexp.MustCompile("^" + rules.regexpr(0) + "$")
			//fmt.Println(zeroRE)
			messages = true
			continue
		}
		if !messages { // rules
			var r rule
			var n int
			a := &(r.l)
			for _, tok := range strings.Split(m, " ") {
				switch {
				case strings.HasSuffix(tok, ":"):
					n = atoi(strings.TrimSuffix(tok, ":"))
				case tok == "|":
					a = &(r.r)
				case tok == `"a"` || tok == `"b"`:
					r.raw = strings.Trim(tok, `"`)
				default:
					*a = append(*a, atoi(tok))
				}
			}
			rules[n] = r
			continue
		}
		// messages
		if zeroRE.MatchString(m) {
			valid++
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(valid)
}
