package mos6510_2

import (
	"log"
)

func (C *CPU) tsx() {
	switch C.Inst.addr {
	case implied:
		C.X = C.SP
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.X)
	C.updateZ(C.X)

}

func (C *CPU) txs() {
	switch C.Inst.addr {
	case implied:
		C.SP = C.X
	default:
		log.Fatal("Bad addressing mode")
	}

}
