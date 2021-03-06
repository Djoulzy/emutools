package mos6510

///////////////////////////////////////////////////////
//                        AND                        //
///////////////////////////////////////////////////////

func (C *CPU) AND_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.A &= C.ram.Read(C.PC)
		C.PC++
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.A &= C.ram.Read(uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_zpx() {
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
		C.A &= C.ram.Read(uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_abs() {
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
		C.A &= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_abx() {
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
			C.Inst.Cycles++
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
			C.A &= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		C.A &= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_aby() {
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
			C.Inst.Cycles++
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
			C.A &= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		C.A &= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_inx() {
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
		C.A &= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_iny() {
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
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.IndAddrLO += C.Y
	case 5:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.A &= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 6:
		C.A &= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) AND_izp() {
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
	case 5:
		C.A &= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        ORA                        //
///////////////////////////////////////////////////////

func (C *CPU) ORA_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.A |= C.ram.Read(C.PC)
		C.PC++
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.A |= C.ram.Read(uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_zpx() {
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
		C.A |= C.ram.Read(uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_abs() {
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
		C.A |= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_abx() {
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
			C.Inst.Cycles++
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
			C.A |= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		C.A |= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_aby() {
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
			C.Inst.Cycles++
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
			C.A |= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		C.A |= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_inx() {
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
		C.A |= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_iny() {
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
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.IndAddrLO += C.Y
	case 5:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.A |= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 6:
		C.A |= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) ORA_izp() {
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
	case 5:
		C.A |= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        EOR                        //
///////////////////////////////////////////////////////

func (C *CPU) EOR_imm() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.A ^= C.ram.Read(C.PC)
		C.PC++
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.A ^= C.ram.Read(uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_zpx() {
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
		C.A ^= C.ram.Read(uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_abs() {
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
		C.A ^= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_abx() {
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
			C.Inst.Cycles++
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
			C.A ^= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		C.A ^= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_aby() {
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
			C.Inst.Cycles++
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
			C.A ^= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 5:
		C.A ^= C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_inx() {
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
		C.A ^= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_iny() {
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
			C.Inst.Cycles++
		} else {
			C.pageCrossed = false
		}
		C.IndAddrLO += C.Y
	case 5:
		tmp := C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))

		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.A ^= tmp
			C.updateN(C.A)
			C.updateZ(C.A)
			C.CycleCount = 0
		}
	case 6:
		C.A ^= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) EOR_izp() {
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
	case 5:
		C.A ^= C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}
