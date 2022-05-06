package mos6510_2

import "log"

func (C *CPU) clc() {
	switch C.Inst.addr {
	case implied:
		C.setC(false)
	default:
		log.Fatal("Bad addressing mode")
	}

}

func (C *CPU) sec() {
	switch C.Inst.addr {
	case implied:
		C.setC(true)
	default:
		log.Fatal("Bad addressing mode")
	}

}
