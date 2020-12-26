package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var (
	pidRE = regexp.MustCompile(`^\d{9}$`)
	eyrRE = regexp.MustCompile(`^20(2\d)|(30)$`)
	iyrRE = regexp.MustCompile(`^20(1\d)|(20)$`)
	hclRE = regexp.MustCompile(`^\#[0-9a-f]{6}$`)
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m, p...)
	os.Exit(1)
}

func main() {
	f, err := os.Open("input.4")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	want := map[string]func(string) bool{
		"ecl": func (s string) bool {
			return map[string]bool{
				"amb": true,
				"blu": true, 
				"brn": true, 
				"gry": true, 
				"grn": true, 
				"hzl": true, 
				"oth": true,
			}[s]
		},
		"pid": pidRE.MatchString,
		"eyr": eyrRE.MatchString,
		"hcl": hclRE.MatchString,
		"byr": func (s string) bool {
			y, err := strconv.Atoi(s)
			if err != nil {
				return false
			}
			return y >= 1920 && y <= 2002
		},
		"iyr": iyrRE.MatchString,
		//"cid":{},
		"hgt": func (s string) bool {
			if len(s) <= 2 {
				return false
			}
			t := strings.TrimSuffix(strings.TrimSuffix(s, "cm"), "in")
			h, err := strconv.Atoi(t)
			if err != nil {
				return false
			}
			if strings.HasSuffix(s, "cm") {
				return h >= 150 && h <= 193
			}
			if strings.HasSuffix(s, "in") {
				return h >= 59 && h <= 76
			}
			return false
		},
	}
	var pp map[string]struct{}
	reset := func() {
		pp = make(map[string]struct{})
		for k := range want {
			pp[k] = struct{}{}
		}
	}
	reset()
	valid := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		m := sc.Text()
		if m == "" {
			if len(pp) == 0 {
				valid++
			}
			reset()
		}
		for _, tok := range strings.Split(m, " ") {
			kv := strings.Split(tok, ":")
			if len(kv) != 2 {
				continue
			}
			k, v := kv[0], kv[1]
			f, ok := want[k]
			if !ok {
				fmt.Printf("%q not found\n", k)
				continue
			} 
			if !f(v) {
				fmt.Printf("%q:%q did not validate\n", k, v)
				continue
			}
			delete(pp, kv[0])
			
		}
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	if len(pp) == 0 {
		valid++
	}
	fmt.Println(valid)
}
