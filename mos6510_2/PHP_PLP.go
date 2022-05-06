package mos6510_2

import (
	"log"
)



func (C *CPU) php() {
	switch C.Inst.addr {
	case implied:
		tmp := C.S
		tmp |= ^B_mask
		tmp |= ^U_mask
		C.pushByteStack(tmp)
	default:
		log.Fatal("Bad addressing mode")
	}

}



func (C *CPU) plp() {
	switch C.Inst.addr {
	case implied:
		C.S = C.pullByteStack()
		C.setB(false)
		C.setU(false)
	default:
		log.Fatal("Bad addressing mode")
	}

}
