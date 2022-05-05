package mos6510_2

import (
	"fmt"
	"log"
)

func (C *CPU) rla() {
	fmt.Printf("%s\nNot implemented: %v\n", C.Trace(), C.Inst)
}

func (C *CPU) rol() {
	var val uint16
	var dest uint16

	switch C.Inst.addr {
	case implied:
		val = uint16(C.A) << 1
		if C.issetC() {
			val++
		}
		C.A = byte(val)
	case zeropage:
		val = uint16(C.ram.Read(C.Oper))
		C.ram.Write(C.Oper, byte(val))
		val <<= 1
		if C.issetC() {
			val++
		}
		C.ram.Write(C.Oper, byte(val))
	case zeropageX:
		dest = C.Oper + uint16(C.X)
		val = uint16(C.ram.Read(dest))
		C.ram.Write(dest, byte(val))
		val <<= 1
		if C.issetC() {
			val++
		}
		C.ram.Write(dest, byte(val))
	case absolute:
		val = uint16(C.ram.Read(C.Oper))
		C.ram.Write(C.Oper, byte(val))
		val <<= 1
		if C.issetC() {
			val++
		}
		C.ram.Write(C.Oper, byte(val))
	case absoluteX:
		dest = C.Oper + uint16(C.X)
		val = uint16(C.ram.Read(dest))
		C.ram.Write(dest, byte(val))
		val <<= 1
		if C.issetC() {
			val++
		}
		C.ram.Write(dest, byte(val))
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(byte(val))
	C.updateZ(byte(val))
	C.setC(val > 0x00FF)

}

func (C *CPU) ror() {
	var val byte

	switch C.Inst.addr {
	case implied:
		carry := C.A&0b00000001 > 0
		C.A >>= 1
		if C.issetC() {
			C.A |= 0b10000000
		}
		C.setC(carry)
		val = C.A
	case zeropage:
		val = C.ram.Read(C.Oper)
		C.ram.Write(C.Oper, val)
		carry := val&0b00000001 > 0
		val >>= 1
		if C.issetC() {
			val |= 0b10000000
		}
		C.setC(carry)
		C.ram.Write(C.Oper, val)
	case zeropageX:
		dest := C.Oper + uint16(C.X)
		val = C.ram.Read(dest)
		C.ram.Write(dest, val)
		carry := val&0b00000001 > 0
		val >>= 1
		if C.issetC() {
			val |= 0b10000000
		}
		C.setC(carry)
		C.ram.Write(dest, val)
	case absolute:
		val = C.ram.Read(C.Oper)
		C.ram.Write(C.Oper, val)
		carry := val&0b00000001 > 0
		val >>= 1
		if C.issetC() {
			val |= 0b10000000
		}
		C.setC(carry)
		C.ram.Write(C.Oper, val)
	case absoluteX:
		dest := C.Oper + uint16(C.X)
		val = C.ram.Read(dest)
		C.ram.Write(dest, val)
		carry := val&0b00000001 > 0
		val >>= 1
		if C.issetC() {
			val |= 0b10000000
		}
		C.setC(carry)
		C.ram.Write(dest, val)
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(byte(val))
	C.updateZ(byte(val))

}

func (C *CPU) sax() {
	fmt.Printf("Not implemented: %v\n", C.Inst)
}

func (C *CPU) slo() {
	fmt.Printf("Not implemented: %v\n", C.Inst)
}

func (C *CPU) sre() {
	fmt.Printf("Not implemented: %v\n", C.Inst)
}
