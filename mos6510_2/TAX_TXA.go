package mos6510_2

import (
	"log"
)

func (C *CPU) tax() {
	switch C.Inst.addr {
	case implied:
		C.X = C.A
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.X)
	C.updateZ(C.X)

}

func (C *CPU) txa() {
	switch C.Inst.addr {
	case implied:
		C.A = C.X
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.A)
	C.updateZ(C.A)

}
