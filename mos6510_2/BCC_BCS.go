package mos6510_2

import "log"

///////////////////////////////////////////////////////
//                        BCC                        //
///////////////////////////////////////////////////////

func (C *CPU) bcc() {
	switch C.Inst.addr {
	case relative:
		C.Oper = C.getRelativeAddr(C.Oper)
		if !C.issetC() {
			C.Inst.addr = Branching
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case Branching:
		if C.PC&0xFF00 == C.Oper&0xFF00 {
			C.PC = C.Oper
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case CrossPage:
		C.PC = C.Oper
	default:
		log.Fatal("Bad addressing mode")
	}
}

func (C *CPU) BCC_rel() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.instCode = C.ram.Read(C.PC)
		if !C.issetC() {
			C.PC += uint16(C.OperLO)
		} else {
			C.PC++
		}
	case 4:
		C.IndAddrLO = C.stack[C.SP]
		C.SP++
	case 5:
		C.IndAddrHI = C.stack[C.SP]
	case 6:
		C.PC = (uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO)
		C.PC++
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        BCS                        //
///////////////////////////////////////////////////////

func (C *CPU) BCS_rel() {}

func (C *CPU) bcs() {
	switch C.Inst.addr {
	case relative:
		C.Oper = C.getRelativeAddr(C.Oper)
		if C.issetC() {
			C.Inst.addr = Branching
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case Branching:
		if C.PC&0xFF00 == C.Oper&0xFF00 {
			C.PC = C.Oper
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case CrossPage:
		C.PC = C.Oper
	default:
		log.Fatal("Bad addressing mode")
	}
}
