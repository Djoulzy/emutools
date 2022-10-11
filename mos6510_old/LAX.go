package mos6510

///////////////////////////////////////////////////////
//                        LAX                        //
///////////////////////////////////////////////////////

func (C *CPU) LAX_zep() {
	switch C.CycleCount {
	case 1:
		C.PC++
	case 2:
		C.OperLO = C.ram.Read(C.PC)
		C.PC++
	case 3:
		C.A = C.ram.Read(uint16(C.OperLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) LAX_zpy() {
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
		C.A = C.ram.Read(uint16(C.OperLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) LAX_abs() {
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
		C.A = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) LAX_aby() {
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
		C.A = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		if C.pageCrossed {
			C.OperHI++
		} else {
			C.CycleCount = 0
		}
	case 5:
		C.A = C.ram.Read((uint16(C.OperHI) << 8) + uint16(C.OperLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) LAX_inx() {
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
		C.A = C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}

func (C *CPU) LAX_iny() {
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
		C.A = C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		if C.pageCrossed {
			C.IndAddrHI++
		} else {
			C.CycleCount = 0
		}
	case 6:
		C.A = C.ram.Read((uint16(C.IndAddrHI) << 8) + uint16(C.IndAddrLO))
		C.X = C.A
		C.updateN(C.A)
		C.updateZ(C.A)
		C.CycleCount = 0
	}
}
