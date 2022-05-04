package mos6510

import (
	"log"
)

func (C *CPU) lda_imm() {
	switch C.cycleCount {
	case 1:
		C.PC++
	case 2:
		C.A = C.ram.Read(C.PC)
		C.PC++
		C.updateN(C.A)
		C.updateZ(C.A)
	}
}

func (C *CPU) lda_zep() {
	switch C.cycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.A = C.ram.Read(uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
	}
}

func (C *CPU) lda_zpx() {
	switch C.cycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperAddr += uint16(C.OperLO + C.X)
	case 4:
		C.A = C.ram.Read(C.OperAddr)
		C.updateN(C.A)
		C.updateZ(C.A)
	}
}

func (C *CPU) lda_abs() {
	switch C.cycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		C.PC++
	case 4:
		C.A = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
	}
}

func (C *CPU) lda_abx() {
	switch C.cycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.X)) > 0x00FF {
			C.pageCrossed = true
		} else {
			C.pageCrossed = false
		} 
		C.OperLO += C.X
		C.PC++
	case 4:
		C.A = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
	}
}
func (C *CPU) lda_aby() {}
func (C *CPU) lda_inx() {}
func (C *CPU) lda_iny() {}

func (C *CPU) lda() {
	var crossed bool

	switch C.Inst.addr {
	case immediate:
		C.A = byte(C.Oper)
	case zeropageX:
		C.A = C.ram.Read(C.Oper + uint16(C.X))
	case zeropage:
		fallthrough
	case absolute:
		C.A = C.ram.Read(C.Oper)
	case absoluteX:
		C.cross_oper = C.Oper + uint16(C.X)
		if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
			C.A = C.ram.Read(C.cross_oper)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case absoluteY:
		C.cross_oper = C.Oper + uint16(C.Y)
		if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
			C.A = C.ram.Read(C.cross_oper)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case indirectX:
		C.A = C.ReadIndirectX(C.Oper)
	case indirectY:
		C.cross_oper = C.GetIndirectYAddr(C.Oper, &crossed)
		if crossed {
			C.A = C.ram.Read(C.cross_oper)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case CrossPage:
		C.A = C.ram.Read(C.cross_oper)
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.A)
	C.updateZ(C.A)
}

func (C *CPU) sta() {
	switch C.Inst.addr {
	case zeropage:
		C.ram.Write(C.Oper, C.A)
	case zeropageX:
		C.ram.Write(C.Oper+uint16(C.X), C.A)
	case absolute:
		C.ram.Write(C.Oper, C.A)
	case absoluteX:
		C.ram.Write(C.Oper+uint16(C.X), C.A)
	case absoluteY:
		C.ram.Write(C.Oper+uint16(C.Y), C.A)
	case indirectX:
		C.WriteIndirectX(C.Oper, C.A)
	case indirectY:
		C.WriteIndirectY(C.Oper, C.A)
	default:
		log.Fatal("Bad addressing mode")
	}

}

func (C *CPU) ldx() {
	switch C.Inst.addr {
	case immediate:
		C.X = byte(C.Oper)
	case zeropage:
		C.X = C.ram.Read(C.Oper)
	case zeropageY:
		C.X = C.ram.Read(C.Oper + uint16(C.Y))
	case absolute:
		C.X = C.ram.Read(C.Oper)
	case absoluteY:
		C.X = C.ram.Read(C.Oper + uint16(C.Y))
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.X)
	C.updateZ(C.X)

}

func (C *CPU) stx() {
	switch C.Inst.addr {
	case zeropage:
		C.ram.Write(C.Oper, C.X)
	case zeropageY:
		C.ram.Write(C.Oper+uint16(C.Y), C.X)
	case absolute:
		C.ram.Write(C.Oper, C.X)
	default:
		log.Fatal("Bad addressing mode")
	}

}

func (C *CPU) ldy() {
	switch C.Inst.addr {
	case immediate:
		C.Y = byte(C.Oper)
	case zeropage:
		C.Y = C.ram.Read(C.Oper)
	case zeropageX:
		C.Y = C.ram.Read(C.Oper + uint16(C.X))
	case absolute:
		C.Y = C.ram.Read(C.Oper)
	case absoluteX:
		C.Y = C.ram.Read(C.Oper + uint16(C.X))
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(C.Y)
	C.updateZ(C.Y)

}

func (C *CPU) sty() {
	switch C.Inst.addr {
	case zeropage:
		C.ram.Write(C.Oper, C.Y)
	case zeropageX:
		C.ram.Write(C.Oper+uint16(C.X), C.Y)
	case absolute:
		C.ram.Write(C.Oper, C.Y)
	default:
		log.Fatal("Bad addressing mode")
	}

}
