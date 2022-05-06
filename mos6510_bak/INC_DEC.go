package mos6510_bak

import (
	"log"
)

func (C *CPU) dec() {
	var val byte

	switch C.Inst.addr {
	case zeropage:
		val = C.ram.Read(C.Oper)
		C.ram.Write(C.Oper, val)
		val--
		C.ram.Write(C.Oper, val)
	case zeropageX:
		oper := C.Oper + uint16(C.X)
		val = C.ram.Read(oper)
		C.ram.Write(oper, val)
		val--
		C.ram.Write(oper, val)
	case absolute:
		val = C.ram.Read(C.Oper)
		C.ram.Write(C.Oper, val)
		val--
		C.ram.Write(C.Oper, val)
	case absoluteX:
		oper := C.Oper + uint16(C.X)
		C.ram.Write(oper, val)
		val--
		C.ram.Write(oper, val)
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(val)
	C.updateZ(val)

}

func (C *CPU) dex() {
	switch C.Inst.addr {
	case implied:
		C.X -= 1
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.X)
	C.updateZ(C.X)

}

func (C *CPU) dey() {
	switch C.Inst.addr {
	case implied:
		C.Y -= 1
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.Y)
	C.updateZ(C.Y)

}

func (C *CPU) inc() {
	var val byte

	switch C.Inst.addr {
	case zeropage:
		val = C.ram.Read(C.Oper)
		C.ram.Write(C.Oper, val)
		val++
		C.ram.Write(C.Oper, val)
	case zeropageX:
		oper := C.Oper + uint16(C.X)
		val = C.ram.Read(oper)
		C.ram.Write(oper, val)
		val++
		C.ram.Write(oper, val)
	case absolute:
		val = C.ram.Read(C.Oper)
		C.ram.Write(C.Oper, val)
		val++
		C.ram.Write(C.Oper, val)
	case absoluteX:
		oper := C.Oper + uint16(C.X)
		C.ram.Write(oper, val)
		val++
		C.ram.Write(oper, val)
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(val)
	C.updateZ(val)
}

func (C *CPU) inx() {
	switch C.Inst.addr {
	case implied:
		C.X += 1
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.X)
	C.updateZ(C.X)

}

func (C *CPU) iny() {
	switch C.Inst.addr {
	case implied:
		C.Y += 1
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.Y)
	C.updateZ(C.Y)

}
