package mos6510

import (
	"fmt"
	"log"
)

func (C *CPU) alr() {
	// var val byte

	log.Fatal("Illegal: ALR")

	// switch C.Inst.addr {
	// case immediate:
	// 	C.A &= byte(C.Oper)
	// 	C.setC(C.A&0x01 == 0x01)
	// 	val = C.A >> 1
	// 	C.A = val
	// default:
	// 	log.Fatal("Bad addressing mode")
	// }
	// C.setN(false)
	// C.updateZ(byte(val))
}

func (C *CPU) sbx() {

	log.Fatal("Illegal: SBX")

	// switch C.Inst.addr {
	// case immediate:
	// 	C.X = (C.A & C.X) - byte(C.Oper)
	// default:
	// 	log.Fatal("Bad addressing mode")
	// }
	// C.setC(C.X >= 0)
	// C.updateN(C.X)
	// C.updateZ(C.X)
}

func (C *CPU) anc() {

	fmt.Printf("Not implemented: %02X\n", C.instCode)

	// switch C.Inst.addr {
	// case immediate:
	// 	C.A &= byte(C.Oper)
	// default:
	// 	log.Fatal("Bad addressing mode")
	// }
	// C.updateN(C.A)
	// C.updateZ(C.A)
	// C.setC(C.A&0b10000000 > 0)
}

func (C *CPU) isc() {
	// var val int
	// var oper byte
	// var crossed bool

	log.Printf("Illegal: ISC - PC: %04X", C.PC)
	log.Fatal("GAME OVER")

	// switch C.Inst.addr {
	// case immediate:
	// 	oper = byte(C.Oper + 1)
	// 	val = int(C.A) - int(oper)
	// 	if C.getC() == 0 {
	// 		val -= 1
	// 	}
	// 	C.updateV(C.A, ^oper, byte(val))
	// 	C.A = byte(val)
	// case zeropage:
	// 	fallthrough
	// case absolute:
	// 	oper = byte(C.ram.Read(C.Oper) + 1)
	// 	val = int(C.A) - int(oper)
	// 	if C.getC() == 0 {
	// 		val -= 1
	// 	}
	// 	C.updateV(C.A, ^oper, byte(val))
	// 	C.A = byte(val)
	// case zeropageX:
	// 	oper = byte(C.ram.Read(C.Oper+uint16(C.X)) + 1)
	// 	val = int(C.A) - int(oper)
	// 	if C.getC() == 0 {
	// 		val -= 1
	// 	}
	// 	C.updateV(C.A, ^oper, byte(val))
	// 	C.A = byte(val)
	// case absoluteX:
	// 	C.cross_oper = C.Oper + uint16(C.X)
	// 	if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
	// 		oper = C.ram.Read(C.cross_oper) + 1
	// 		val = int(C.A) - int(oper)
	// 		if C.getC() == 0 {
	// 			val -= 1
	// 		}
	// 		C.updateV(C.A, ^oper, byte(val))
	// 		C.A = byte(val)
	// 	} else {
	// 		C.Inst.addr = CrossPage
	// 		// C.State = Compute
	// 		C.Inst.Cycles++
	// 		return
	// 	}
	// case absoluteY:
	// 	C.cross_oper = C.Oper + uint16(C.Y)
	// 	if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
	// 		oper = C.ram.Read(C.cross_oper) + 1
	// 		val = int(C.A) - int(oper)
	// 		if C.getC() == 0 {
	// 			val -= 1
	// 		}
	// 		C.updateV(C.A, ^oper, byte(val))
	// 		C.A = byte(val)
	// 	} else {
	// 		C.Inst.addr = CrossPage
	// 		// C.State = Compute
	// 		C.Inst.Cycles++
	// 		return
	// 	}
	// case indirectX:
	// 	oper = byte(C.ReadIndirectX(C.Oper) + 1)
	// 	val = int(C.A) - int(oper)
	// 	if C.getC() == 0 {
	// 		val -= 1
	// 	}
	// 	C.updateV(C.A, ^oper, byte(val))
	// 	C.A = byte(val)
	// case indirectY:
	// 	C.cross_oper = C.GetIndirectYAddr(C.Oper, &crossed)
	// 	if crossed {
	// 		oper = C.ram.Read(C.cross_oper) + 1
	// 		val = int(C.A) - int(oper)
	// 		if C.getC() == 0 {
	// 			val -= 1
	// 		}
	// 		C.updateV(C.A, ^oper, byte(val))
	// 		C.A = byte(val)
	// 	} else {
	// 		C.Inst.addr = CrossPage
	// 		// C.State = Compute
	// 		C.Inst.Cycles++
	// 		return
	// 	}
	// case CrossPage:
	// 	oper = C.ram.Read(C.cross_oper) + 1
	// 	val = int(C.A) - int(oper)
	// 	if C.getC() == 0 {
	// 		val -= 1
	// 	}
	// 	C.updateV(C.A, ^oper, byte(val))
	// 	C.A = byte(val)
	// default:
	// 	log.Fatal("Bad addressing mode")
	// }
	// C.setC(val >= 0x00)
	// C.setN(val&0b10000000 == 0b10000000)
	// C.updateZ(byte(val))
}

func (C *CPU) dcp() {
	// var val int
	// var crossed bool

	log.Fatal("Illegal: DCP")

	// switch C.Inst.addr {
	// case immediate:
	// 	val = int(C.A) - int(C.Oper-1)
	// case zeropage:
	// 	val = int(C.A) - int(C.ram.Read(C.Oper)-1)
	// case zeropageX:
	// 	val = int(C.A) - int(C.ram.Read(C.Oper+uint16(C.X))-1)
	// case absolute:
	// 	val = int(C.A) - int(C.ram.Read(C.Oper)-1)
	// case absoluteX:
	// 	C.cross_oper = C.Oper + uint16(C.X)
	// 	if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
	// 		val = int(C.A) - int(C.ram.Read(C.cross_oper)) - 1
	// 	} else {
	// 		C.Inst.addr = CrossPage
	// 		// C.State = Compute
	// 		C.Inst.Cycles++
	// 		return
	// 	}
	// case absoluteY:
	// 	C.cross_oper = C.Oper + uint16(C.Y)
	// 	if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
	// 		val = int(C.A) - int(C.ram.Read(C.cross_oper)) - 1
	// 	} else {
	// 		C.Inst.addr = CrossPage
	// 		// C.State = Compute
	// 		C.Inst.Cycles++
	// 		return
	// 	}
	// case indirectX:
	// 	val = int(C.A) - int(C.ReadIndirectX(C.Oper)) - 1
	// case indirectY:
	// 	C.cross_oper = C.GetIndirectYAddr(C.Oper, &crossed)
	// 	if crossed {
	// 		val = int(C.A) - int(C.ram.Read(C.cross_oper)) - 1
	// 	} else {
	// 		C.Inst.addr = CrossPage
	// 		// C.State = Compute
	// 		C.Inst.Cycles++
	// 		return
	// 	}
	// case CrossPage:
	// 	val = int(C.A) - int(C.ram.Read(C.cross_oper)) - 1
	// default:
	// 	log.Fatal("Bad addressing mode")
	// }
	// C.setC(val >= 0)
	// C.updateN(byte(val))
	// C.updateZ(byte(val))
}

func (C *CPU) rla() {
	fmt.Printf("Not implemented: %02X\n", C.instCode)
}

func (C *CPU) sax() {
	fmt.Printf("Not implemented: %02X\n", C.instCode)
}

func (C *CPU) slo() {
	fmt.Printf("Not implemented: %02X\n", C.instCode)
}

func (C *CPU) sre() {
	fmt.Printf("Not implemented: %02X\n", C.instCode)
}
