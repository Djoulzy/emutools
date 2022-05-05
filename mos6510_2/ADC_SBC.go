package mos6510_2

import (
	"log"

	"github.com/albenik/bcd"
)

///////////////////////////////////////////////////////
//                        ADC                        //
///////////////////////////////////////////////////////

func (C *CPU) doADC(inputVal byte) byte {
	if C.getD() == 1 {
		val_bcd := uint16(bcd.ToUint8(C.A)) + uint16(bcd.ToUint8(inputVal)) + uint16(C.getC())
		vals := bcd.FromUint16(val_bcd)
		C.setC(vals[0] > 0)
		C.setV(val_bcd > 100)
		return vals[1]
	} else {
		val := uint16(C.A) + uint16(inputVal) + uint16(C.getC())
		C.setC(val > 0x00FF)
		C.updateV(C.A, inputVal, byte(val))
		return byte(val)
	}
}

func (C *CPU) ADC_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		tmp := C.ram.Read(C.PC)
		C.PC++
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		tmp := C.ram.Read(uint16(C.OperLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_zpx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.OperLO = C.OperLO + C.X
	case 4:
		tmp := C.ram.Read(uint16(C.OperLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_abs() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		C.PC++
	case 4:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_abx() {
	switch C.CycleCount {
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
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.A = C.doADC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doADC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_aby() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.OperHI = C.ram.Read(C.PC)
		if (uint16(C.OperLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
		} else {
			C.pageCrossed = false
		}
		C.OperLO += C.Y
		C.PC++
	case 4:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.A = C.doADC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		tmp := C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.A = C.doADC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_inx() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.ram.Read(uint16(C.OperLO))
		C.Pointer = C.OperLO + C.X
	case 4:
		C.IndAddrLO = C.ram.Read(uint16(C.Pointer))
	case 5:
		C.IndAddrHI = C.ram.Read(uint16(C.Pointer + 1))
	case 6:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.A = C.doADC(tmp)

		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ADC_iny() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.IndAddrLO = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.IndAddrHI = C.ram.Read(uint16(C.OperLO + 1))
		if (uint16(C.IndAddrLO) + uint16(C.Y)) > 0x00FF {
			C.pageCrossed = true
		} else {
			C.pageCrossed = false
		}
		C.IndAddrLO += C.Y
	case 5:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))

		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.A = C.doADC(tmp)
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 6:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.A = C.doADC(tmp)
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) adc() {
	var val uint16
	var oper byte
	var crossed bool

	switch C.Inst.addr {
	case immediate:
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(byte(C.Oper)))
			val_bcd := A_bcd + oper_bcd + uint16(C.getC())
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[0] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = uint16(C.A) + C.Oper + uint16(C.getC())
			C.setC(val > 0x00FF)
			C.updateV(C.A, byte(C.Oper), byte(val))
			C.A = byte(val)
		}
	case zeropage:
		fallthrough
	case absolute:
		oper = C.ram.Read(C.Oper)
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(oper))
			val_bcd := A_bcd + oper_bcd + uint16(C.getC())
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[0] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = uint16(C.A) + uint16(oper) + uint16(C.getC())
		}
		C.setC(val > 0x00FF)
		C.updateV(C.A, oper, byte(val))
		C.A = byte(val)
	case zeropageX:
		oper = C.ram.Read(C.Oper + uint16(C.X))
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(oper))
			val_bcd := A_bcd + oper_bcd + uint16(C.getC())
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[0] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = uint16(C.A) + uint16(oper) + uint16(C.getC())
		}
		C.setC(val > 0x00FF)
		C.updateV(C.A, oper, byte(val))
		C.A = byte(val)
	case absoluteX:
		C.cross_oper = C.Oper + uint16(C.X)
		if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
			oper = C.ram.Read(C.cross_oper)
			if C.getD() == 1 {
				A_bcd := uint16(bcd.ToUint8(C.A))
				oper_bcd := uint16(bcd.ToUint8(oper))
				val_bcd := A_bcd + oper_bcd + uint16(C.getC())
				vals := bcd.FromUint16(val_bcd)
				C.setC(vals[0] > 0)
				C.setV(val_bcd > 100)
				C.updateN(vals[1])
				C.updateZ(vals[1])
				C.A = vals[1]
				return
			} else {
				val = uint16(C.A) + uint16(oper) + uint16(C.getC())
			}
			C.setC(val > 0x00FF)
			C.updateV(C.A, oper, byte(val))
			C.A = byte(val)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case absoluteY:
		C.cross_oper = C.Oper + uint16(C.Y)
		if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
			oper = C.ram.Read(C.cross_oper)
			if C.getD() == 1 {
				A_bcd := uint16(bcd.ToUint8(C.A))
				oper_bcd := uint16(bcd.ToUint8(oper))
				val_bcd := A_bcd + oper_bcd + uint16(C.getC())
				vals := bcd.FromUint16(val_bcd)
				C.setC(vals[0] > 0)
				C.setV(val_bcd > 100)
				C.updateN(vals[1])
				C.updateZ(vals[1])
				C.A = vals[1]
				return
			} else {
				val = uint16(C.A) + uint16(oper) + uint16(C.getC())
			}
			C.setC(val > 0x00FF)
			C.updateV(C.A, oper, byte(val))
			C.A = byte(val)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case indirectX:
		oper = C.ReadIndirectX(C.Oper)
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(oper))
			val_bcd := A_bcd + oper_bcd + uint16(C.getC())
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[0] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = uint16(C.A) + uint16(oper) + uint16(C.getC())
		}
		C.setC(val > 0x00FF)
		C.updateV(C.A, oper, byte(val))
		C.A = byte(val)
	case indirectY:
		C.cross_oper = C.GetIndirectYAddr(C.Oper, &crossed)
		if crossed {
			oper = C.ram.Read(C.cross_oper)
			if C.getD() == 1 {
				A_bcd := uint16(bcd.ToUint8(C.A))
				oper_bcd := uint16(bcd.ToUint8(oper))
				val_bcd := A_bcd + oper_bcd + uint16(C.getC())
				vals := bcd.FromUint16(val_bcd)
				C.setC(vals[0] > 0)
				C.setV(val_bcd > 100)
				C.updateN(vals[1])
				C.updateZ(vals[1])
				C.A = vals[1]
				return
			} else {
				val = uint16(C.A) + uint16(oper) + uint16(C.getC())
			}
			C.setC(val > 0x00FF)
			C.updateV(C.A, oper, byte(val))
			C.A = byte(val)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case CrossPage:
		oper = C.ram.Read(C.cross_oper)
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(oper))
			val_bcd := A_bcd + oper_bcd + uint16(C.getC())
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[0] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = uint16(C.A) + uint16(oper) + uint16(C.getC())
		}
		C.setC(val > 0x00FF)
		C.updateV(C.A, oper, byte(val))
		C.A = byte(val)
	default:
		log.Fatal("Bad addressing mode")
	}
	C.updateN(byte(val))
	C.updateZ(byte(val))
}

///////////////////////////////////////////////////////
//                        SBC                        //
///////////////////////////////////////////////////////

func (C *CPU) sbc() {
	var val int
	var oper byte
	var crossed bool

	switch C.Inst.addr {
	case immediate:
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(byte(C.Oper)))
			val_bcd := A_bcd - oper_bcd
			if C.getC() == 0 {
				val_bcd -= 1
			}
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[1] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = int(C.A) - int(C.Oper)
		}
		if C.getC() == 0 {
			val -= 1
		}
		C.updateV(C.A, ^byte(C.Oper), byte(val))
		C.A = byte(val)
	case zeropage:
		fallthrough
	case absolute:
		oper = C.ram.Read(C.Oper)
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(byte(oper)))
			val_bcd := A_bcd - oper_bcd
			if C.getC() == 0 {
				val_bcd -= 1
			}
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[1] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = int(C.A) - int(oper)
		}
		if C.getC() == 0 {
			val -= 1
		}
		C.updateV(C.A, ^oper, byte(val))
		C.A = byte(val)
	case zeropageX:
		oper = C.ram.Read(C.Oper + uint16(C.X))
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(byte(oper)))
			val_bcd := A_bcd - oper_bcd
			if C.getC() == 0 {
				val_bcd -= 1
			}
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[1] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = int(C.A) - int(oper)
		}
		if C.getC() == 0 {
			val -= 1
		}
		C.updateV(C.A, ^oper, byte(val))
		C.A = byte(val)
	case absoluteX:
		C.cross_oper = C.Oper + uint16(C.X)
		if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
			oper = C.ram.Read(C.cross_oper)
			if C.getD() == 1 {
				A_bcd := uint16(bcd.ToUint8(C.A))
				oper_bcd := uint16(bcd.ToUint8(byte(oper)))
				val_bcd := A_bcd - oper_bcd
				if C.getC() == 0 {
					val_bcd -= 1
				}
				vals := bcd.FromUint16(val_bcd)
				C.setC(vals[1] > 0)
				C.setV(val_bcd > 100)
				C.updateN(vals[1])
				C.updateZ(vals[1])
				C.A = vals[1]
				return
			} else {
				val = int(C.A) - int(oper)
			}
			if C.getC() == 0 {
				val -= 1
			}
			C.updateV(C.A, ^oper, byte(val))
			C.A = byte(val)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case absoluteY:
		C.cross_oper = C.Oper + uint16(C.Y)
		if C.Oper&0xFF00 == C.cross_oper&0xFF00 {
			oper = C.ram.Read(C.cross_oper)
			if C.getD() == 1 {
				A_bcd := uint16(bcd.ToUint8(C.A))
				oper_bcd := uint16(bcd.ToUint8(byte(oper)))
				val_bcd := A_bcd - oper_bcd
				if C.getC() == 0 {
					val_bcd -= 1
				}
				vals := bcd.FromUint16(val_bcd)
				C.setC(vals[1] > 0)
				C.setV(val_bcd > 100)
				C.updateN(vals[1])
				C.updateZ(vals[1])
				C.A = vals[1]
				return
			} else {
				val = int(C.A) - int(oper)
			}
			if C.getC() == 0 {
				val -= 1
			}
			C.updateV(C.A, ^oper, byte(val))
			C.A = byte(val)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case indirectX:
		oper = C.ReadIndirectX(C.Oper)
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(byte(oper)))
			val_bcd := A_bcd - oper_bcd
			if C.getC() == 0 {
				val_bcd -= 1
			}
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[1] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = int(C.A) - int(oper)
		}
		if C.getC() == 0 {
			val -= 1
		}
		C.updateV(C.A, ^oper, byte(val))
		C.A = byte(val)
	case indirectY:
		C.cross_oper = C.GetIndirectYAddr(C.Oper, &crossed)
		if crossed {
			oper = C.ram.Read(C.cross_oper)
			if C.getD() == 1 {
				A_bcd := uint16(bcd.ToUint8(C.A))
				oper_bcd := uint16(bcd.ToUint8(byte(oper)))
				val_bcd := A_bcd - oper_bcd
				if C.getC() == 0 {
					val_bcd -= 1
				}
				vals := bcd.FromUint16(val_bcd)
				C.setC(vals[1] > 0)
				C.setV(val_bcd > 100)
				C.updateN(vals[1])
				C.updateZ(vals[1])
				C.A = vals[1]
				return
			} else {
				val = int(C.A) - int(oper)
			}
			if C.getC() == 0 {
				val -= 1
			}
			C.updateV(C.A, ^oper, byte(val))
			C.A = byte(val)
		} else {
			C.Inst.addr = CrossPage
			C.State = Compute
			C.Inst.Cycles++
			return
		}
	case CrossPage:
		oper = C.ram.Read(C.cross_oper)
		if C.getD() == 1 {
			A_bcd := uint16(bcd.ToUint8(C.A))
			oper_bcd := uint16(bcd.ToUint8(byte(oper)))
			val_bcd := A_bcd - oper_bcd
			if C.getC() == 0 {
				val_bcd -= 1
			}
			vals := bcd.FromUint16(val_bcd)
			C.setC(vals[1] > 0)
			C.setV(val_bcd > 100)
			C.updateN(vals[1])
			C.updateZ(vals[1])
			C.A = vals[1]
			return
		} else {
			val = int(C.A) - int(oper)
		}
		if C.getC() == 0 {
			val -= 1
		}
		C.updateV(C.A, ^oper, byte(val))
		C.A = byte(val)
	default:
		log.Fatal("Bad addressing mode")
	}
	C.setC(val >= 0x00)
	C.setN(val&0b10000000 == 0b10000000)
	C.updateZ(byte(val))
}
