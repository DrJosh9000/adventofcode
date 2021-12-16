package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/bearmini/bitstream-go"
)

func main() {
	f, err := os.Open("inputs/16.txt")
	if err != nil {
		log.Fatalf("Couldn't open: %v", err)
	}
	defer f.Close()
	_, p, err := parsePacket(bitstream.NewReader(hex.NewDecoder(f), nil))
	if err != nil {
		log.Fatalf("Couldn't parse packet: %v")
	}

	fmt.Println(eval(p))
}

func parsePacket(r *bitstream.Reader) (int, packet, error) {
	ver, err := r.ReadNBitsAsUint8(3)
	if err != nil {
		return 0, nil, err
	}
	typ, err := r.ReadNBitsAsUint8(3)
	if err != nil {
		return 0, nil, err
	}
	hdr := header{ver, typ}
	switch typ {
	case 4:
		// literal
		groups, value := 0, 0
		for {
			group, err := r.ReadNBitsAsUint8(5)
			if err != nil {
				return 0, nil, err
			}
			value <<= 4
			value += int(group & 0b00001111)
			groups++
			if group&0b00010000 == 0 {
				// final group
				break
			}
		}
		return 5*groups + 6, &literal{hdr, value}, nil
	default:
		// operator
		o := &operator{hdr, nil}
		ltyp, err := r.ReadBool()
		if err != nil {
			return 0, nil, err
		}
		if ltyp {
			// subpacket count
			count, err := r.ReadNBitsAsUint16BE(11)
			if err != nil {
				return 0, nil, err
			}
			bits := 6 + 1 + 11
			o.subpackets = make([]packet, 0, count)
			for i := uint16(0); i < count; i++ {
				n, p, err := parsePacket(r)
				if err != nil {
					return 0, nil, err
				}
				o.subpackets = append(o.subpackets, p)
				bits += n
			}
			return bits, o, nil
		} else {
			// total bit length
			length, err := r.ReadNBitsAsUint16BE(15)
			if err != nil {
				return 0, nil, err
			}
			bits := 6 + 1 + 15 + int(length)
			read := 0
			for {
				n, p, err := parsePacket(r)
				if err != nil {
					return 0, nil, err
				}
				o.subpackets = append(o.subpackets, p)
				read += n
				if read >= int(length) {
					break
				}
			}
			return bits, o, nil
		}
	}
}

func eval(p packet) int {
	if l, ok := p.(*literal); ok {
		return l.value
	}
	o := p.(*operator)
	switch o.PacketType() {
	case 0:
		s := 0
		for _, q := range o.subpackets {
			s += eval(q)
		}
		return s
	case 1:
		s := 1
		for _, q := range o.subpackets {
			s *= eval(q)
		}
		return s
	case 2:
		m := math.MaxInt
		for _, q := range o.subpackets {
			if x := eval(q); x < m {
				m = x
			}
		}
		return m
	case 3:
		m := math.MinInt
		for _, q := range o.subpackets {
			if x := eval(q); x > m {
				m = x
			}
		}
		return m
	case 4:
		log.Fatal("Operator with packet type 4?")
	case 5:
		if eval(o.subpackets[0]) > eval(o.subpackets[1]) {
			return 1
		}
		return 0
	case 6:
		if eval(o.subpackets[0]) < eval(o.subpackets[1]) {
			return 1
		}
		return 0
	case 7:
		if eval(o.subpackets[0]) == eval(o.subpackets[1]) {
			return 1
		}
		return 0
	}
	return -666
}

type packet interface {
	Version() byte
	PacketType() byte
}

type header struct {
	version, ptype byte
}

func (h header) Version() byte    { return h.version }
func (h header) PacketType() byte { return h.ptype }

type literal struct {
	header
	value int
}

func (l *literal) String() string { return fmt.Sprintf("lit{v%d %d}", l.version, l.value) }

type operator struct {
	header
	subpackets []packet
}

func (o *operator) String() string {
	return fmt.Sprintf("op{v%d %d %v}", o.version, o.ptype, o.subpackets)
}
