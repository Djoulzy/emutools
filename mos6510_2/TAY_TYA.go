package mos6510_2

import (
	"log"
)

func (C *CPU) tay() {
	switch C.Inst.addr {
	case implied:
		C.Y = C.A
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.Y)
	C.updateZ(C.Y)

}

func (C *CPU) tya() {
	switch C.Inst.addr {
	case implied:
		C.A = C.Y
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.A)
	C.updateZ(C.A)

}
