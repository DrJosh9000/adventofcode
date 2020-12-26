package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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

func eval(x ast.Expr) int {
	switch e := x.(type) {
	case *ast.BasicLit:
		if e.Kind != token.INT {
			die("Literal must be type INT, got %v", e.Kind)
		}
		return atoi(e.Value)
	case *ast.ParenExpr:
		return eval(e.X)
	case *ast.BinaryExpr:
		switch e.Op {
		case token.ADD:
			// hax
			return eval(e.X) * eval(e.Y)
		case token.MUL:
			return eval(e.X) + eval(e.Y)
		}
	default:
		die("Unknown expression element of type %T", x)
	}
	return 0
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
		m = strings.Map(func(r rune) rune {
			switch r {
			case '*':
				return '+'
			case '+':
				return '*'
			default:
				return r
			}
		}, m)
		x, err := parser.ParseExpr(m)
		if err != nil {
			die("Couldn't parse %q: %v", m, err)
		}
		sum += eval(x)
	}
	if err := sc.Err(); err != nil {
		die("Couldn't read file: %v", err)
	}
	fmt.Println(sum)
}
