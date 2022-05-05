package mos6510_2

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
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff--
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
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
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff--
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
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
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff--
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
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
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff--
	case 7:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
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
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 4:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff++
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
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
		C.RMWBuff = C.ram.Read(uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff++
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
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
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 5:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff++
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
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
		C.RMWBuff = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
	case 6:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)
		C.setC(C.RMWBuff&0b10000000 > 1)
		C.RMWBuff++
	case 7:
		C.ram.Write(uint16(C.OperLO), C.RMWBuff)

		C.updateN(C.RMWBuff)
		C.updateZ(C.RMWBuff)
		C.CycleCount = 0
	}
}
