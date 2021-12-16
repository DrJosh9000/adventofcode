package main

import (
	"encoding/hex"
	"fmt"
	"log"
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

	//fmt.Println(p)

	var sum func(packet) int
	sum = func(p packet) int {
		switch x := p.(type) {
		case *literal:
			return int(x.Version())
		case *operator:
			s := int(x.Version())
			for _, q := range x.subpackets {
				s += sum(q)
			}
			return s
		}
		return -666
	}
	fmt.Println(sum(p))
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
		groups := 0
		for {
			group, err := r.ReadNBitsAsUint8(5)
			if err != nil {
				return 0, nil, err
			}
			groups++
			if group&0b00010000 == 0 {
				// final group
				break
			}
		}
		return 5*groups + 6, &literal{hdr, 0}, nil
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
