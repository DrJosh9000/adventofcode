package main

import (
	"fmt"
	"io"
	"os"
)

func die(m string, p ...interface{}) {
	fmt.Fprintf(os.Stderr, m+"\n", p...)
	os.Exit(1)
}

func abs(x int) int {
	if x<0 {
		return -x
	}
	return x
}

func main() {
	f, err := os.Open("input.12")
	if err != nil {
		die("Couldn't open file: %v", err)
	}
	defer f.Close()

	type ins struct {
		opcode rune
		operand int
	}
	var x, y int
	wx, wy := 10, 1
	rot := func(t int) {
		for t<0 {t+=360}
		for t>=360 {t-=360}
		switch t {
		case 90:
			wx, wy = -wy, wx
		case 180:
			wx, wy = -wx, -wy
		case 270:
			wx, wy = wy, -wx
		}
	}
	for {
		in := &ins{}
		_, err := fmt.Fscanf(f, "%c%d\n", &in.opcode, &in.operand)
		if err == io.EOF {
			break
		}		
		if err != nil {
			die("Couldn't fscanf: %v", err)
		}
		switch in.opcode {
		case 'N':
			wy+=in.operand
		case 'S':
			wy-=in.operand
		case 'E':
			wx+=in.operand
		case 'W':
			wx-=in.operand
		case 'L':
			rot(in.operand)
		case 'R':
			rot(-in.operand)
		case 'F':
			x+=wx*in.operand
			y+=wy*in.operand
		} 
	}
	fmt.Println(abs(x)+abs(y))
}
