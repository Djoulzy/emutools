package mos6510_2

import "log"

func (C *CPU) cli() {
	switch C.Inst.addr {
	case implied:
		C.setI(false)
	default:
		log.Fatal("Bad addressing mode")
	}

}

func (C *CPU) sei() {
	switch C.Inst.addr {
	case implied:
		C.setI(true)
	default:
		log.Fatal("Bad addressing mode")
	}

}
