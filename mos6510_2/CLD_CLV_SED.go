package mos6510_2

import (
	"log"
)



func (C *CPU) cld() {
	switch C.Inst.addr {
	case implied:
		C.setD(false)
	default:
		log.Fatal("Bad addressing mode")
	}

}



func (C *CPU) clv() {
	switch C.Inst.addr {
	case implied:
		C.setV(false)
	default:
		log.Fatal("Bad addressing mode")
	}

}



func (C *CPU) sed() {
	// log.Fatal("Decimal mode")
	switch C.Inst.addr {
	case implied:
		C.setD(true)
	default:
		log.Fatal("Bad addressing mode")
	}

}

