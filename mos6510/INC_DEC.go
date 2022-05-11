package mos6510

///////////////////////////////////////////////////////
//                        DEC                        //
///////////////////////////////////////////////////////

func (C *CPU) DEC_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff--
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) DEC_zpx() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff--
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) DEC_abs() {
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
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff--
	case 6:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) DEC_abx() {
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
		C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		}
	case 5:
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff--
	case 7:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

///////////////////////////////////////////////////////
//                        INC                        //
///////////////////////////////////////////////////////

func (C *CPU) INC_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff++
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) INC_zpx() {
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
		C.tmpBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff++
	case 6:
		C.ram.Write(uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) INC_abs() {
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
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff++
	case 6:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}

func (C *CPU) INC_abx() {
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
		C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		if C.pageCrossed {
			C.OperHI++
		}
	case 5:
		C.tmpBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)
		C.setC(C.tmpBuff&0b10000000 > 1)
		C.tmpBuff++
	case 7:
		C.ram.Write((uint16(C.OperHI) << 8) + uint16(C.OperLO), C.tmpBuff)

		C.updateN(C.tmpBuff)
		C.updateZ(C.tmpBuff)
		C.CycleCount = 0
	}
}
