package mos6510_2

import "log"

func (C *CPU) pha() {
	switch C.Inst.addr {
	case implied:
		C.pushByteStack(C.A)
	default:
		log.Fatal("Bad addressing mode")
	}

}

func (C *CPU) pla() {
	switch C.Inst.addr {
	case implied:
		C.A = C.pullByteStack()
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.A)
	C.updateZ(C.A)

}
